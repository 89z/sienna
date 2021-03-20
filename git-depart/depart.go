package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func run(name string, arg ...string) error {
   c := exec.Command(name, arg...)
   c.Stderr = os.Stderr
   c.Stdout = os.Stdout
   x.LogInfo("Run", name, arg)
   return c.Run()
}

func main() {
   e := run("git", "commit", "--verbose")
   check(e)
   _, e = os.Stat("config.toml")
   if e != nil {
      return
   }
   e = os.RemoveAll("docs")
   check(e)
   e = run("hugo")
   check(e)
   e = run("git", "add", ".")
   check(e)
   e = run("git", "commit", "--amend", "--no-edit")
   check(e)
}
