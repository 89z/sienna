package main

import (
   "fmt"
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
   dirs, err := os.ReadDir(root)
   if err != nil {
      panic(err)
   }
   for _, each := range dirs {
      cmd := exec.Command(name, arg...)
      cmd.Dir = filepath.Join(root, each.Name())
      cmd.Stdout = os.Stdout
      fmt.Println("\x1b[7m Dir \x1b[m", cmd.Dir)
      fmt.Println("\x1b[7m Run \x1b[m", cmd)
      err := cmd.Run()
      if err != nil {
         panic(err)
      }
   }
}
