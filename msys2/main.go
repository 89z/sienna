package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "os"
   "path"
)

func main() {
   // both
   if len(os.Args) != 3 {
      println(`msys2 query mingw-w64-x86_64-gcc
msys2 sync gcc.txt`)
      os.Exit(1)
   }
   cache, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }
   var inst install
   for _, each := range []string{
      "mingw/x86_64/mingw64.db.tar.gz", "msys/x86_64/msys.db.tar.gz",
   } {
      mirror.Path = each
      inst = newInstall(mirror, cache, cache, "sienna", "msys2")
      _, e = x.Copy(inst.source, inst.cache)
      if e == nil {
         x.LogInfo("Gz", inst.cache)
         var tar extract.Archive
         tar.Gz(inst.cache, inst.dest)
      } else if os.IsExist(e) {
         x.LogInfo("Exist", inst.cache)
      } else {
         log.Fatal(e)
      }
   }
   dirs, e := os.ReadDir(inst.dest)
   if e != nil {
      log.Fatal(e)
   }
   db := newDatabase()
   for _, each := range dirs {
      desc := path.Join(inst.dest, each.Name(), "desc")
      db.scan(desc)
   }
   target := os.Args[2]
   switch os.Args[1] {
   case "query":
      db.query(target)
   case "sync":
   }
}
