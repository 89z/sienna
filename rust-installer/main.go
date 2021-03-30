package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "log"
   "os"
   "path"
)

func main() {
   var (
      cache userCache
      err error
   )
   cache.dir, err = os.UserCacheDir()
   if err != nil {
      log.Fatal(err)
   }
   cache.dir = path.Join(cache.dir, "sienna", "rust")
   var dist distChannel
   err = cache.unmarshal(&dist)
   if err != nil {
      log.Fatal(err)
   }
   for _, each := range []target{
      dist.Pkg.Cargo, dist.Pkg.RustStd, dist.Pkg.Rustc,
   } {
      err = cache.install(each.Target.X8664PcWindowsGnu.XzUrl)
      if err != nil {
         log.Fatal(err)
      }
   }
}
