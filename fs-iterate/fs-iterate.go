package main

import (
   "github.com/89z/x"
   "io/ioutil"
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
   x.Check(e)
   for _, each := range dirs {
      dir := filepath.Join(root, each.Name())
      println(x.ColorCyan(dir))
      cmd := exec.Command(name, arg...)
      cmd.Dir = dir
      cmd.Stdout = os.Stdout
      e = cmd.Run()
      x.Check(e)
   }
}
