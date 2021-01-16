package main

import (
   "log"
   "os"
)

func main() {
   if len(os.Args) != 3 {
      println(`synopsis:
   msys2 <operation> <target>

examples:
   msys2 deps mingw-w64-x86_64-libgit2
   msys2 sync git.txt`)
      os.Exit(1)
   }

   oper_s := os.Args[1]
   tar_s := os.Args[2]

   o, e := newManager()
   if e != nil {
      log.Fatal(e)
   }

   if oper_s == "deps" {
      m, e := o.Resolve(tar_s)
      if e != nil {
         log.Fatal(e)
      }
      for s := range m {
         println(s)
      }
   } else {
      e := o.Sync(tar_s)
      if e != nil {
         log.Fatal(e)
      }
   }
}
