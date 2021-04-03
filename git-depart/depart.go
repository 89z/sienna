package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
)

func run(name string, arg ...string) error {
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   x.LogInfo("Run", name, arg)
   return c.Run()
}

func depart() error {
   e := run("git", "commit", "--verbose")
   if e != nil { return e }
   _, e = os.Stat("config.toml")
   if e != nil { return nil }
   e = os.RemoveAll("docs")
   if e != nil { return e }
   e = run("hugo")
   if e != nil { return e }
   e = run("git", "add", ".")
   if e != nil { return e }
   return run("git", "commit", "--amend", "--no-edit")
}

func main() {
   e := depart()
   if e != nil {
      log.Fatal(e)
   }
}
