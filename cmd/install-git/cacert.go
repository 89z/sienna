package main

import (
   "net/http"
   "os"
   "path/filepath"
)

func main() {
   addr := "https://curl.haxx.se/ca/cacert.pem"
   println("Get", addr)
   res, err := http.Get(addr)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := `C:\sienna\msys2\usr\ssl\certs`
   os.MkdirAll(dst, os.ModeDir)
   file, err := os.Create(filepath.Join(dst, "ca-bundle.crt"))
   if err != nil {
      panic(err)
   }
   defer file.Close()
   file.ReadFrom(res.Body)
}
