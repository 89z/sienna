package main

import (
   "github.com/89z/x"
   "log"
   "os"
)

func main() {
   inst := x.NewInstall("sienna/msys2/usr/ssl/certs", "ca-bundle.crt")
   os.Remove(inst.Dest)
   _, e := x.Copy("https://curl.haxx.se/ca/cacert.pem", inst.Dest)
   if e != nil {
      log.Fatal(e)
   }
}
