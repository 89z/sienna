package main

import (
   "fmt"
   "github.com/mholt/archiver/v3"
   "github.com/pelletier/go-toml"
   "log"
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

func main() {
   cache_s, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }

   toml_s, e := GetDatabase(cache_s)
   if e != nil {
      log.Fatal(e)
   }

   toml_o, e := toml.LoadFile(toml_s)
   if e != nil {
      log.Fatal(e)
   }

   a := []string{"cargo", "rust-std", "rustc"}
   for n := range a {
      key_s := "pkg." + a[n] + ".target.x86_64-pc-windows-gnu.xz_url"
      url_s := toml_o.Get(key_s).(string)
      file_s := path.Base(url_s)
      if ! IsFile(file_s) {
         e := Copy(url_s, file_s)
         if e != nil {
            log.Fatal(e)
         }
      }
      e = Unarchive(file_s, `C:\Rust`)
      if e != nil {
         log.Fatal(e)
      }
   }
}
