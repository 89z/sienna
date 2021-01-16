package main

import (
   "github.com/pelletier/go-toml"
   "os"
   "path"
)

func main() {
   cache, e := os.UserCacheDir()
   check(e)
   db, e := getDatabase(cache)
   check(e)
   file, e := toml.LoadFile(db)
   check(e)
   a := []string{"cargo", "rust-std", "rustc"}
   for n := range a {
      key_s := "pkg." + a[n] + ".target.x86_64-pc-windows-gnu.xz_url"
      url_s := file.Get(key_s).(string)
      file_s := path.Base(url_s)
      if ! IsFile(file_s) {
         e := Copy(url_s, file_s)
         check(e)
      }
      e = Unarchive(file_s, `C:\Rust`)
      check(e)
   }
}
