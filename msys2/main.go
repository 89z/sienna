package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "os"
   "path"
)

func main() {
   if len(os.Args) != 3 {
      println(`msys2 query mingw-w64-x86_64-gcc
msys2 sync gcc.txt`)
      os.Exit(1)
   }
   cache, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }
   cache = path.Join(cache, "sienna")
   var tar extract.Tar
   for _, each := range []string{
      "/mingw/x86_64/mingw64.db.tar.gz", "/msys/x86_64/msys.db.tar.gz",
   } {
      archive := path.Join(cache, each)
      _, e = x.Copy(mirror + each, archive)
      if os.IsExist(e) {
         continue
      } else if e != nil {
         log.Fatal(e)
      }
      x.LogInfo("Gz", each)
      e = tar.Gz(
         archive, path.Dir(archive),
      )
      if e != nil {
         log.Fatal(e)
      }
   }
}
