package main

import (
   "github.com/89z/sienna"
   "github.com/pelletier/go-toml"
   "net/http"
   "os"
   "path/filepath"
)

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "rust")
   create := filepath.Join(cache, "channel-rust-stable.toml")
   _, err = os.Stat(create)
   if err != nil {
      get := "https://static.rust-lang.org/dist/channel-rust-stable.toml"
      println("Get", get)
      r, err := http.Get(get)
      if err != nil {
         panic(err)
      }
      defer r.Body.Close()
      f, err := os.Create(create)
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(r.Body)
   } else {
      println("Exist", create)
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
      get := dist.Pkg[pkg].Target["x86_64-pc-windows-gnu"].XZ_URL
      create := filepath.Join(cache, filepath.Base(get))
      _, err := os.Stat(create)
      if err != nil {
         println("Get", get)
         r, err := http.Get(get)
         if err != nil {
            panic(err)
         }
         defer r.Body.Close()
         f, err := os.Create(create)
         if err != nil {
            panic(err)
         }
         defer f.Close()
         f.ReadFrom(r.Body)
      } else {
         println("Exist", create)
      }
      tar := sienna.Archive{2}
      println("Xz", create)
      tar.Xz(create, `C:\sienna\rust`)
   }
}
