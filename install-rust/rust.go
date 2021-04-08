package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

const channel = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

func main() {
   inst := x.NewInstall("sienna/rust", channel)
   inst.SetCache()
   _, e := x.Copy(channel, inst.Cache)
   if os.IsExist(e) {
      x.LogInfo("Exist", inst.Cache)
   } else if e != nil {
      log.Fatal(e)
   }
   cache, e := os.Open(inst.Cache)
   if e != nil {
      log.Fatal(e)
   }
   defer cache.Close()
   var dist struct {
      Pkg map[string]struct {
         Target map[string]struct { XZ_URL string }
      }
   }
   toml.NewDecoder(cache).Decode(&dist)
   for _, each := range []string{"cargo", "rust-std", "rustc"} {
      addr := dist.Pkg[each].Target["x86_64-pc-windows-gnu"].XZ_URL
      inst = x.NewInstall("sienna/rust", addr)
      inst.SetCache()
      _, e = x.Copy(addr, inst.Cache)
      if os.IsExist(e) {
         x.LogInfo("Exist", inst.Cache)
      } else if e != nil {
         log.Fatal(e)
      }
      tar := extract.Archive{2}
      x.LogInfo("Xz", inst.Cache)
      tar.Xz(inst.Cache, inst.Dest)
   }
}
