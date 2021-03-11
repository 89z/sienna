package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "os"
   "os/exec"
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
      cmd := exec.Command(name, arg...)
      x.LogInfo("Dir", dir)
      cmd.Dir = dir
      cmd.Stderr = os.Stderr
      cmd.Stdout = os.Stdout
      e = cmd.Run()
      if e != nil {
         log.Fatal(e)
      }
   }
}
