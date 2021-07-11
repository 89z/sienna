package main

import (
   "fmt"
   "os"
   "os/exec"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func run(name string, arg ...string) error {
   cmd := exec.Command(name, arg...)
   cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
   fmt.Println(invert, "Run", reset, cmd)
   return cmd.Run()
}

func main() {
   err := run("git", "commit", "--verbose")
   if err != nil {
      panic(err)
   }
   if _, err := os.Stat("config.toml"); err != nil {
      // if this not exist, return
      return
   }
   if err := run("hugo"); err != nil {
      panic(err)
   }
   if err := run("git", "add", "."); err != nil {
      panic(err)
   }
   run("git", "commit", "--amend", "--no-edit")
}
