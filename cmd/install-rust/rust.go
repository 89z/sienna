package main

import (
   "github.com/89z/rosso"
   "github.com/pelletier/go-toml"
   "os"
)

const channel = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

func main() {
   inst := rosso.NewInstall("sienna/rust", channel)
   inst.SetCache()
   _, e := rosso.Copy(channel, inst.Cache)
   if os.IsExist(e) {
      rosso.LogInfo("Exist", inst.Cache)
   } else if e != nil {
      panic(e)
   }
   cache, e := os.Open(inst.Cache)
   if e != nil {
      panic(e)
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
      inst = rosso.NewInstall("sienna/rust", addr)
      inst.SetCache()
      _, e = rosso.Copy(addr, inst.Cache)
      if os.IsExist(e) {
         rosso.LogInfo("Exist", inst.Cache)
      } else if e != nil {
         panic(e)
      }
      tar := rosso.Archive{2}
      rosso.LogInfo("Xz", inst.Cache)
      tar.Xz(inst.Cache, inst.Dest)
   }
}
