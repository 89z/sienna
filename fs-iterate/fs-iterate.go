package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "os"
)

func main() {
   if len(os.Args) < 3 {
      println("fs-iterate <path> <command>")
      os.Exit(1)
   }
   root, name, arg := os.Args[1], os.Args[2], os.Args[3:]
   e := os.Chdir(root)
   x.Check(e)
   dirs, e := ioutil.ReadDir(".")
   x.Check(e)
   for _, dir := range dirs {
      path := dir.Name()
      println(x.ColorCyan(path))
      e = os.Chdir(path)
      x.Check(e)
      e = x.System(name, arg...)
      x.Check(e)
      e = os.Chdir("..")
      x.Check(e)
   }
}
