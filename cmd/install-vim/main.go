package main

import (
   "fmt"
   "github.com/89z/sienna"
   "net/http"
   "os"
   "path/filepath"
)

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "vim")
   create := filepath.Join(cache, version)
   if err := getCreate(
      "https://github.com/vim/vim-win32-installer/releases/download/" + version,
      create,
   ); err != nil {
      panic(err)
   }
   arc := sienna.Archive{2}
   fmt.Println("Zip", create)
   arc.Zip(create, `C:\sienna\vim`)
   for _, rt := range runtime {
      get := "https://raw.githubusercontent.com/" + rt.get + rt.create
      fmt.Println("Get", get)
      r, err := http.Get(get)
      if err != nil {
         panic(err)
      }
      defer r.Body.Close()
      f, err := os.Create(filepath.Join(`C:\sienna\vim`, rt.create))
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(r.Body)
   }
}
