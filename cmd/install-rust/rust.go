package main

import (
   "github.com/89z/sienna"
   "github.com/pelletier/go-toml"
   "net/http"
   "os"
   "path/filepath"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func getCreate(get, create string) error {
   _, err := os.Stat(create)
   if err == nil {
      println(invert, "Exist", reset, create)
      return nil
   }
   println(invert, "Get", reset, get)
   res, err := http.Get(get)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(create)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}


func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "rust")
   create := filepath.Join(cache, "channel-rust-stable.toml")
   get := "https://static.rust-lang.org/dist/channel-rust-stable.toml"
   if err := getCreate(get, create); err != nil {
      panic(err)
   }
   file, err := os.Open(create)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   var dist struct {
      Pkg map[string]struct {
         Target map[string]struct {
            XZ_URL string
         }
      }
   }
   toml.NewDecoder(file).Decode(&dist)
   for _, pkg := range []string{"cargo", "rust-std", "rustc"} {
      create := filepath.Join(cache, filepath.Base(get))
      get := dist.Pkg[pkg].Target["x86_64-pc-windows-gnu"].XZ_URL
      err := getCreate(get, create)
      if err != nil {
         panic(err)
      }
      tar := sienna.Archive{2}
      println(invert, "Xz", reset, create)
      tar.Xz(create, `C:\sienna\rust`)
   }
}
