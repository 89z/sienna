package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "os"
)

func main() {
   if len(os.Args) < 3 {
      println("fs-iterate <path> <command>")
      os.Exit(1)
   }
   root, name, arg := os.Args[1], os.Args[2], os.Args[3:]
   os.Chdir(root)
   dirs, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for _, dir := range dirs {
      path := dir.Name()
      println(x.ColorCyan(path))
      os.Chdir(path)
      e = x.System(name, arg...)
      if e != nil {
         log.Fatal(e)
      }
      os.Chdir("..")
   }
}
