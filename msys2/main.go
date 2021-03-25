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
   //db := newDatabase()
   for _, each := range []pack{
      {cache, "mingw", "x86_64", "mingw64.db.tar.gz"},
      {cache, "msys", "x86_64", "msys.db.tar.gz"},
   } {
      _, e = x.Copy(each.remote(), each.local())
      if e == nil {
         x.LogInfo("Gz", each.local())
         tar.Gz(each.local(), each.dir())
      } else if os.IsExist(e) {
         x.LogInfo("Exist", each.local())
      } else {
         log.Fatal(e)
      }
      dirs, e := os.ReadDir(each.dir())
      if e != nil {
         log.Fatal(e)
      }
      for _, each := range dirs {
         println(each)
      }
   }
}
