package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "os"
   "path/filepath"
)

func main() {
   if len(os.Args) < 3 {
      println("fs-iterate <path> <command>")
      os.Exit(1)
   }
   root, name, arg := os.Args[1], os.Args[2], os.Args[3:]
   dirs, e := ioutil.ReadDir(root)
   if e != nil {
      log.Fatal(e)
   }
   for _, each := range dirs {
      dir := filepath.Join(root, each.Name())
      cmd := x.Command(name, arg...)
      cmd.Dir = dir
      e = cmd.Run()
      if e != nil {
         log.Fatal(e)
      }
   }
}
