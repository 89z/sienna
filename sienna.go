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
   if err != nil {
      return err
   }
   defer file.Close()
   rXz, err := xz.NewReader(file)
   if err != nil {
      return err
   }
   return a.tarCreate(rXz, dest)
}

func (a Archive) Zip(source, dest string) error {
   rZip, err := zip.OpenReader(source)
   if err != nil {
      return err
   }
   defer rZip.Close()
   for _, entry := range rZip.File {
      if entry.Mode().IsDir() {
         continue
      }
      name := a.strip(dest, entry.Name)
      if name == "" {
         continue
      }
      os.MkdirAll(filepath.Dir(name), os.ModeDir)
      open, err := entry.Open()
      if err != nil {
         return err
      }
      file, err := os.Create(name)
      if err != nil {
         return err
      }
      defer file.Close()
      file.ReadFrom(open)
   }
   return nil
}

func (a Archive) Zst(source, dest string) error {
   file, err := os.Open(source)
   if err != nil {
      return err
   }
   defer file.Close()
   rZst, err := zstd.NewReader(file)
   if err != nil {
      return err
   }
   return a.tarCreate(rZst, dest)
}

func (a Archive) strip(left, right string) string {
   split := strings.SplitN(right, "/", a.Strip + 1)
   if len(split) <= a.Strip {
      return ""
   }
   return filepath.Join(left, split[a.Strip])
}

func (a Archive) tarCreate(source io.Reader, dest string) error {
   rTar := tar.NewReader(source)
   for {
      cur, err := rTar.Next()
      if err == io.EOF {
         break
      } else if err != nil {
         return err
      }
      name := a.strip(dest, cur.Name)
      if name == "" {
         continue
      }
      switch cur.Typeflag {
      case tar.TypeLink:
         _, err := os.Stat(name)
         if err == nil {
            os.Remove(name)
         }
         os.Link(a.strip(dest, cur.Linkname), name)
      case tar.TypeReg:
         os.MkdirAll(filepath.Dir(name), os.ModeDir)
         create, err := os.Create(name)
         if err != nil {
            return err
         }
         defer create.Close()
         create.ReadFrom(rTar)
      }
   }
   return nil
}
