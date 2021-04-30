package main

import (
   "github.com/89z/rosso"
   "os"
)

func main() {
   inst := rosso.NewInstall("sienna/msys2/usr/ssl/certs", "ca-bundle.crt")
   os.Remove(inst.Dest)
   _, e := rosso.Copy("https://curl.haxx.se/ca/cacert.pem", inst.Dest)
   if e != nil {
      panic(e)
   }
}
