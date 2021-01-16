package main

import (
   "fmt"
   "github.com/mholt/archiver/v3"
   "log"
   "os"
   "path"
   "sienna"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func getDatabase(cache_s string) (string, error) {
   rust := path.Join(cache_s, "Rust")
   os.Mkdir(rust, os.ModeDir)
   os.Chdir(rust)
   url := "https://static.rust-lang.org/dist/channel-rust-stable.toml"
   file := path.Base(url)
   if sienna.IsFile(file) {
      return file, nil
   }
   _, e := sienna.HttpCopy(url, file)
   if e != nil {
      return "", e
   }
   return file, nil
}

func unarchive(file_s, dir_s string) error {
   tar_o := &archiver.Tar{OverwriteExisting: true, StripComponents: 2}
   fmt.Println("EXTRACT", file_s)
   xz_o := archiver.TarXz{Tar: tar_o}
   return xz_o.Unarchive(file_s, dir_s)
}
