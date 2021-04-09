package main

import (
   "github.com/89z/x"
   "log"
   "os"
)

func depart() error {
   var c x.Cmd
   e := c.Run("git", "commit", "--verbose")
   if e != nil { return e }
   _, e = os.Stat("config.toml")
   if e != nil { return nil }
   e = os.RemoveAll("docs")
   if e != nil { return e }
   e = c.Run("hugo")
   if e != nil { return e }
   e = c.Run("git", "add", ".")
   if e != nil { return e }
   return c.Run("git", "commit", "--amend", "--no-edit")
}

func main() {
   e := depart()
   if e != nil {
      log.Fatal(e)
   }
}
