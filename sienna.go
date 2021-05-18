package sienna

import (
   "archive/tar"
   "archive/zip"
   "compress/bzip2"
   "compress/gzip"
   "github.com/klauspost/compress/zstd"
   "github.com/xi2/xz"
   "io"
   "os"
   "path/filepath"
   "strings"
   "testing/fstest"
)

func TarGzMemory(source string) (fstest.MapFS, error) {
   file, err := os.Open(source)
   if err != nil { return nil, err }
   defer file.Close()
   gzRead, err := gzip.NewReader(file)
   if err != nil { return nil, err }
   tarRead := tar.NewReader(gzRead)
   files := make(fstest.MapFS)
   for {
      cur, err := tarRead.Next()
      if err == io.EOF { break } else if err != nil { return nil, err }
      if cur.Typeflag != tar.TypeReg { continue }
      data, err := io.ReadAll(tarRead)
      if err != nil { return nil, err }
      files[cur.Name] = &fstest.MapFile{Data: data}
   }
   return files, nil
}

type Archive struct { Strip int }

func (a Archive) Bz2(source, dest string) error {
   open, err := os.Open(source)
   if err != nil { return err }
   return a.tarCreate(bzip2.NewReader(open), dest)
}

func (a Archive) Gz(source, dest string) error {
   open, err := os.Open(source)
   if err != nil { return err }
   gzRead, err := gzip.NewReader(open)
   if err != nil { return err }
   return a.tarCreate(gzRead, dest)
}

func (a Archive) Xz(source, dest string) error {
   open, err := os.Open(source)
   if err != nil { return err }
   xzRead, err := xz.NewReader(open, 0)
   if err != nil { return err }
   return a.tarCreate(xzRead, dest)
}

func (a Archive) Zip(source, dest string) error {
   read, err := zip.OpenReader(source)
   if err != nil { return err }
   for _, file := range read.File {
      if file.Mode().IsDir() { continue }
      name := a.strip(dest, file.Name)
      if name == "" { continue }
      err = os.MkdirAll(filepath.Dir(name), os.ModeDir)
      if err != nil { return err }
      open, err := file.Open()
      if err != nil { return err }
      create, err := os.Create(name)
      if err != nil { return err }
      defer create.Close()
      _, err = create.ReadFrom(open)
      if err != nil { return err }
   }
   return nil
}

func (a Archive) Zst(source, dest string) error {
   open, err := os.Open(source)
   if err != nil { return err }
   zstRead, err := zstd.NewReader(open)
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
         _, err = os.Stat(name)
         if err == nil {
            err = os.Remove(name)
            if err != nil {
               return err
            }
         }
         err = os.Link(a.strip(dest, cur.Linkname), name)
         if err != nil { return err }
      case tar.TypeReg:
         err = os.MkdirAll(filepath.Dir(name), os.ModeDir)
         if err != nil { return err }
         create, err := os.Create(name)
         if err != nil { return err }
         defer create.Close()
         _, err = create.ReadFrom(tarRead)
         if err != nil { return err }
      }
   }
   return nil
}
