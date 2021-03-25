package main

import (
   "fmt"
   "os"
   "testing"
)

func TestInstall(t *testing.T) {
   cache, e := os.UserCacheDir()
   if e != nil {
      t.Error(e)
   }
   var inst install
   // example 1
   mirror.Path = "msys/x86_64/msys.db.tar.gz"
   inst = newInstall(mirror, cache, cache, "sienna", "msys2")
   fmt.Println(inst)
   // example 2
   mirror.Path = "msys/x86_64/zstd-1.4.8-1-x86_64.pkg.tar.zst"
   inst = newInstall(mirror, cache, `C:\`, "sienna", "msys2")
   fmt.Println(inst)
}
