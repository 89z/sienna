package main

import (
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "os"
   "path/filepath"
)

func main() {
   user, e := os.UserCacheDir()
   x.Check(e)
   e = os.Chdir(filepath.Join(user, "rust"))
   x.Check(e)
   dist := filepath.Base(channel)
   if ! x.IsFile(dist) {
      _, e = x.Copy(channel, dist)
      x.Check(e)
   }
   manifest, e := toml.LoadFile(dist)
   x.Check(e)
   for _, pack := range packages {
      url := manifest.M(pack).S("xz_url")
      base := filepath.Base(url)
      if ! x.IsFile(base) {
         _, e = x.Copy(url, base)
         x.Check(e)
      }
      e = unarchive(base, `C:\rust`)
      x.Check(e)
   }
}
