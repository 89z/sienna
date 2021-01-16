package main

import (
   "github.com/pelletier/go-toml"
   "os"
   "path"
   "sienna"
)

func main() {
   cache, e := os.UserCacheDir()
   check(e)
   db, e := getDatabase(cache)
   check(e)
   file, e := toml.LoadFile(db)
   check(e)
   for _, pack := range []string{"cargo", "rust-std", "rustc"} {
      key := "pkg." + pack + ".target.x86_64-pc-windows-gnu.xz_url"
      source := file.Get(key).(string)
      dest := path.Base(source)
      if ! sienna.IsFile(dest) {
         _, e := sienna.HttpCopy(source, dest)
         check(e)
      }
      e = unarchive(dest, `C:\Rust`)
      check(e)
   }
}
