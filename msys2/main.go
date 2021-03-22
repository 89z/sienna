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
   cache = path.Join(cache, "sienna", "msys2")
   var tar extract.Tar
   for _, each := range []string{
      "mingw/x86_64/mingw64.db.tar.gz", "msys/x86_64/msys.db.tar.gz",
   } {
      mirror.Path = each
      archive := path.Join(cache, each)
      _, e = x.Copy(
         mirror.String(), archive,
      )
      if e == nil {
         x.LogInfo("Gz", archive)
         e = tar.Gz(
            archive, path.Dir(archive),
         )
         if e != nil {
            log.Fatal(e)
         }
      } else if os.IsExist(e) {
         x.LogInfo("Exist", archive)
      } else {
         log.Fatal(e)
      }
   }
}
