package main

import (
   "github.com/89z/x"
   "os"
   "path/filepath"
)

func main() {
   if len(os.Args) < 3 {
      println("fs-iterate <path> <command>")
      os.Exit(1)
   }
   root, name, arg := os.Args[1], os.Args[2], os.Args[3:]
   dirs, err := os.ReadDir(root)
   if err != nil {
      panic(err)
   }
   for _, each := range dirs {
      dir := filepath.Join(root, each.Name())
      cmd := x.Cmd{dir}
      x.LogInfo("Dir", dir)
      err = cmd.Run(name, arg...)
      if err != nil {
         panic(err)
      }
   }
}
