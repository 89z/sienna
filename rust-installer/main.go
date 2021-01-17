package main

import (
   "os"
   "sienna"
)

func main() {
   cache, e := os.UserCacheDir()
   check(e)
   rust := sienna.NewPath(cache, "rust")
   e = os.Chdir(rust.Join)
   check(e)
   dist := sienna.NewPath(channel)
   if ! sienna.IsFile(dist.Base) {
      _, e = sienna.HttpCopy(channel, dist.Base)
      check(e)
   }
   file, e := sienna.TomlGetFile(dist.Base)
   check(e)
   for _, pack := range []string{"cargo", "rust-std", "rustc"} {
      key := file.S("pkg." + pack + ".target.x86_64-pc-windows-gnu.xz_url")
      xz := sienna.NewPath(key)
      if ! sienna.IsFile(xz.Base) {
         _, e := sienna.HttpCopy(xz.Join, xz.Base)
         check(e)
      }
      e = unarchive(xz.Base, `C:\rust`)
      check(e)
   }
}
