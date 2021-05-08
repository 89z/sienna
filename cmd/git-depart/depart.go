package main

import (
   "fmt"
   "os"
   "os/exec"
)

func run(name string, arg ...string) error {
   cmd := exec.Command(name, arg...)
   cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
   fmt.Println("\x1b[7m Run \x1b[m", cmd)
   return cmd.Run()
}

func main() {
   err := run("git", "commit", "--verbose")
   if err != nil {
      panic(err)
   }
   _, err = os.Stat("config.toml")
   // if this not exist, return
   if err != nil { return }
   os.RemoveAll("docs")
   err = run("hugo")
   if err != nil {
      panic(err)
   }
   err = run("git", "add", ".")
   if err != nil {
      panic(err)
   }
   run("git", "commit", "--amend", "--no-edit")
}
