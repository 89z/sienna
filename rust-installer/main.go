package main

import (
   "github.com/89z/sienna"
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "os"
)

func main() {
   cache, e := os.UserCacheDir()
   check(e)
   rust := sienna.NewPath(cache, "rust")
   e = os.Chdir(rust.Join)
   check(e)
   dist := sienna.NewPath(channel)
   if ! x.IsFile(dist.Base) {
      _, e = x.HttpCopy(channel, dist.Base)
      check(e)
   }
   file, e := toml.LoadFile(dist.Base)
   check(e)
   for _, pack := range []string{"cargo", "rust-std", "rustc"} {
      key := file.S("pkg." + pack + ".target.x86_64-pc-windows-gnu.xz_url")
      xz := sienna.NewPath(key)
      if ! x.IsFile(xz.Base) {
         _, e := x.HttpCopy(xz.Join, xz.Base)
         check(e)
      }
      e = unarchive(xz.Base, `C:\rust`)
      check(e)
   }
}
