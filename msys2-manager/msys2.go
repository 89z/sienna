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

   oper := os.Args[1]
   target := os.Args[2]

   o, e := newManager()
   if e != nil {
      log.Fatal(e)
   }

   if oper == "deps" {
      m, e := o.resolve(target)
      if e != nil {
         log.Fatal(e)
      }
      for s := range m {
         println(s)
      }
   } else {
      e := o.sync(target)
      if e != nil {
         log.Fatal(e)
      }
   }
}
