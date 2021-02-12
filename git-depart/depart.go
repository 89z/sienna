package main

import (
   "log"
   "os"
   "os/exec"
)

func hugo(config string) error {
   _, e := os.Stat(config)
   if e != nil {
      return nil
   }
   e = os.RemoveAll("docs")
   if e != nil {
      return e
   }
   e = exec.Command("hugo").Run()
   if e != nil {
      return e
   }
   e = exec.Command("git", "add", ".").Run()
   if e != nil {
      return e
   }
   return exec.Command("git", "commit", "--amend").Run()
}

func main() {
   e := exec.Command("git", "commit", "--verbose").Run()
   if e != nil {
      log.Fatal(e)
   }
   e = hugo("config.toml")
   if e != nil {
      log.Fatal(e)
   }
}
