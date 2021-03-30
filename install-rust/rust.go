package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

const remote = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

func main() {
   inst, e := x.NewInstall(remote, "sienna", "rust")
   if e != nil {
      log.Fatal(e)
   }
   _, e = x.Copy(remote, inst.Cache)
   if os.IsExist(e) {
      x.LogInfo("Exist", inst.Cache)
   } else if e != nil {
      log.Fatal(e)
   }
   channel, e := os.Open(inst.Cache)
   if e != nil {
      log.Fatal(e)
   }
   defer channel.Close()
   var dist struct {
      Pkg map[string]struct {
         Target map[string]struct { XZ_URL string }
      }
   }
   toml.NewDecoder(channel).Decode(&dist)
   for _, each := range []string{"cargo", "rust-std", "rustc"} {
      source := dist.Pkg[each].Target["x86_64-pc-windows-gnu"].XZ_URL
      inst, e = x.NewInstall(source, "sienna", "rust")
      if e != nil {
         log.Fatal(e)
      }
      _, e = x.Copy(source, inst.Cache)
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
