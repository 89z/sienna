package main

import (
   "github.com/89z/rosso"
   "os"
)

func main() {
   if len(os.Args) != 3 {
      println(`install-msys2 query git
install-msys2 sync git.txt`)
      os.Exit(1)
   }
   db := newDatabase()
   for _, each := range []string{
      "/mingw/ucrt64/ucrt64.db",
      "/mingw/x86_64/mingw64.db",
      "/msys/x86_64/msys.db",
   } {
      inst := rosso.NewInstall("sienna/msys2", each)
      inst.SetCache()
      _, err := rosso.Copy(mirror + each, inst.Cache)
      if os.IsExist(err) {
         rosso.LogInfo("Exist", inst.Cache)
      } else if err != nil {
         panic(err)
      }
      fs, err := rosso.TarGzMemory(inst.Cache)
      if err != nil {
         panic(err)
      }
      for _, each := range fs {
         db.scan(each.Data)
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
