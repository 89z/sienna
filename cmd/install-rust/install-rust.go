package main

import (
   "github.com/89z/rosso"
   "github.com/pelletier/go-toml"
   "os"
   "path/filepath"
)

func main() {
   // get
   get := "https://static.rust-lang.org/dist/channel-rust-stable.toml"
   // cache
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "rust")
   // create
   create := filepath.Join(cache, "channel-rust-stable.toml")
   // copy
   err = rosso.Copy(get, create)
   if os.IsExist(err) {
      println("Exist", create)
   } else if err != nil {
      panic(err)
   }
   file, err := os.Open(create)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   var dist struct {
      Pkg map[string]struct {
         Target map[string]struct { XZ_URL string }
      }
   }
   toml.NewDecoder(file).Decode(&dist)
   for _, pkg := range []string{"cargo", "rust-std", "rustc"} {
      // get
      get := dist.Pkg[pkg].Target["x86_64-pc-windows-gnu"].XZ_URL
      // create
      create := filepath.Join(cache, filepath.Base(get))
      // copy
      err := rosso.Copy(get, create)
      if os.IsExist(err) {
         println("Exist", create)
      } else if err != nil {
         panic(err)
      }
      tar := rosso.Archive{2}
      println("Xz", create)
      tar.Xz(create, `C:\sienna\rust`)
   }
}
