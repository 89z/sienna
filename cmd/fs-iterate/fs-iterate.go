package main

import (
   "fmt"
   "os"
   "os/exec"
   "path/filepath"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func main() {
   if len(os.Args) < 3 {
      fmt.Println("fs-iterate <path> <command>")
      return
   }
   root, name, arg := os.Args[1], os.Args[2], os.Args[3:]
   dirs, err := os.ReadDir(root)
   if err != nil {
      panic(err)
   }
   for _, each := range dirs {
      cmd := exec.Command(name, arg...)
      cmd.Dir = filepath.Join(root, each.Name())
      cmd.Stdout = os.Stdout
      fmt.Println(invert, "Dir", reset, cmd.Dir)
      fmt.Println(invert, "Run", reset, cmd)
      err := cmd.Run()
      if err != nil {
         panic(err)
      }
   }
}
