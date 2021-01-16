package main

import (
   "fmt"
   "github.com/mholt/archiver/v3"
   "os"
   "path"
)

func GetDatabase(cache_s string) (string, error) {
   rust_s := path.Join(cache_s, "Rust")
   os.Mkdir(rust_s, os.ModeDir)
   os.Chdir(rust_s)
   url_s := "https://static.rust-lang.org/dist/channel-rust-stable.toml"

   file_s := path.Base(url_s)
   if IsFile(file_s) {
      return file_s, nil
   }

   e := Copy(url_s, file_s)
   if e != nil {
      return "", e
   }

   return file_s, nil
}

func Unarchive(file_s, dir_s string) error {
   tar_o := &archiver.Tar{OverwriteExisting: true, StripComponents: 2}
   fmt.Println("EXTRACT", file_s)
   xz_o := archiver.TarXz{Tar: tar_o}
   return xz_o.Unarchive(file_s, dir_s)
}
