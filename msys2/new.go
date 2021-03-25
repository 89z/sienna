package main

import (
   "fmt"
   "log"
   "os"
   "path/filepath"
)

type install struct {
   cache string
   destination string
}

func newInstall(source, cache, dest string, base ...string) install {
   for _, each := range base {
      cache = filepath.Join(cache, each)
      dest = filepath.Join(dest, each)
   }
   cache = filepath.Join(cache, filepath.Base(source))
   return install{cache, dest}
}

func main() {
   cache, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }
   var inst install
   inst = newInstall(
      "http://repo.msys2.org/msys/x86_64/msys.db.tar.gz",
      cache,
      cache,
      "sienna",
      "msys2",
   )
   fmt.Println(inst)
   inst = newInstall(
      "http://repo.msys2.org/msys/x86_64/zstd-1.4.8-1-x86_64.pkg.tar.zst",
      cache,
      `C:\`,
      "sienna",
      "msys2",
   )
   fmt.Println(inst)
}
