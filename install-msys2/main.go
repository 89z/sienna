package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "os"
)

func main() {
   if len(os.Args) != 3 {
      println(`msys2 query mingw-w64-x86_64-gcc
msys2 sync gcc.txt`)
      os.Exit(1)
   }
   db := newDatabase()
   for _, each := range []string{
      "mingw/x86_64/mingw64.db.tar.gz", "msys/x86_64/msys.db.tar.gz",
   } {
      mirror.Path = each
      inst := x.NewInstall("sienna/msys2", each)
      inst.SetCache()
      _, e := x.Copy(mirror.String(), inst.Cache)
      if os.IsExist(e) {
         x.LogInfo("Exist", inst.Cache)
      } else if e != nil {
         log.Fatal(e)
      }
      files, e := extract.TarGzMemory(inst.Cache)
      if e != nil {
         log.Fatal(e)
      }
      for _, each := range files {
         db.scan(each)
      }
   }
   target := os.Args[2]
   switch os.Args[1] {
   case "query":
      db.query(target)
   case "sync":
      db.sync(target)
   }
}
