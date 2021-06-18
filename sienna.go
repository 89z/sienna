package sienna

import (
   "archive/tar"
   "archive/zip"
   "github.com/klauspost/compress/zstd"
   "github.com/89z/xz"
   "io"
   "os"
   "path/filepath"
   "strings"
)

type Archive struct {
   Strip int
}

func (a Archive) Xz(source, dest string) error {
   file, err := os.Open(source)
   if err != nil { return err }
   defer file.Close()
   xzRead, err := xz.NewReader(file)
   if err != nil { return err }
   return a.tarCreate(xzRead, dest)
}

func (a Archive) Zip(source, dest string) error {
   read, err := zip.OpenReader(source)
   if err != nil { return err }
   defer read.Close()
   for _, file := range read.File {
      if file.Mode().IsDir() { continue }
      name := a.strip(dest, file.Name)
      if name == "" { continue }
      if err := os.MkdirAll(filepath.Dir(name), os.ModeDir); err != nil {
         return err
      }
      open, err := file.Open()
      if err != nil { return err }
      create, err := os.Create(name)
      if err != nil { return err }
      defer create.Close()
      if _, err := create.ReadFrom(open); err != nil { return err }
   }
   return nil
}

func (a Archive) Zst(source, dest string) error {
   file, err := os.Open(source)
   if err != nil { return err }
   defer file.Close()
   zstRead, err := zstd.NewReader(file)
   if err != nil { return err }
   return a.tarCreate(zstRead, dest)
}

func (a Archive) strip(left, right string) string {
   split := strings.SplitN(right, "/", a.Strip + 1)
   if len(split) <= a.Strip { return "" }
   return filepath.Join(left, split[a.Strip])
}

func (a Archive) tarCreate(source io.Reader, dest string) error {
   tarRead := tar.NewReader(source)
   for {
      cur, err := tarRead.Next()
      if err == io.EOF { break } else if err != nil { return err }
      name := a.strip(dest, cur.Name)
      if name == "" { continue }
      switch cur.Typeflag {
      case tar.TypeLink:
         _, err := os.Stat(name)
         if err == nil {
            os.Remove(name)
         }
         if err := os.Link(a.strip(dest, cur.Linkname), name); err != nil {
            return err
         }
      case tar.TypeReg:
         os.MkdirAll(filepath.Dir(name), os.ModeDir)
         create, err := os.Create(name)
         if err != nil { return err }
         defer create.Close()
         create.ReadFrom(tarRead)
      }
   }
   return nil
}