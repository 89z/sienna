package main

import (
   "net/http"
   "os"
   "path/filepath"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func main() {
   src := "https://curl.haxx.se/ca/cacert.pem"
   println(invert, "Get", reset, src)
   res, err := http.Get(src)
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
