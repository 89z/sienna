package main

import (
   "io/ioutil"
   "log"
   "os"
   "os/exec"
)

func cyan(s string) string {
   return "\x1b[1;36m" + s + "\x1b[m"
}

func system(command ...string) error {
   name, arg := command[0], command[1:]
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   return c.Run()
}

func main() {
   if len(os.Args) < 3 {
      println("fs-iterate <path> <command>")
      os.Exit(1)
   }
   root, command := os.Args[1], os.Args[2:]
   dirs, e := ioutil.ReadDir(root)
   if e != nil {
      log.Fatal(e)
   }
   for _, dir := range dirs {
      name := dir.Name()
      println(cyan(name))
      os.Chdir(name)
      e = system(command...)
      if e != nil {
         log.Fatal(e)
      }
      os.Chdir("..")
   }
}
