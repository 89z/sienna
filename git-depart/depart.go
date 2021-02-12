package main

import (
   "github.com/89z/x"
   "log"
   "os"
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
   e = x.Command("hugo").Run()
   if e != nil {
      return e
   }
   e = x.Command("git", "add", ".").Run()
   if e != nil {
      return e
   }
   return x.Command("git", "commit", "--amend").Run()
}

func main() {
   e := x.Command("git", "commit", "--verbose").Run()
   if e != nil {
      log.Fatal(e)
   }
   e = hugo("config.toml")
   if e != nil {
      log.Fatal(e)
   }
}
