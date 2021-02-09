package main

import (
   "github.com/89z/x"
   "os"
)

func main() {
   e := x.Command("git", "commit", "--verbose").Run()
   x.Check(e)
   _, e = os.Stat("config.toml")
   if e != nil {
      return
   }
   println(x.ColorCyan("Remove"), "docs")
   e = os.RemoveAll("docs")
   x.Check(e)
   e = x.Command("hugo").Run()
   x.Check(e)
   e = x.Command("git", "add", ".").Run()
   x.Check(e)
   e = x.Command("git", "commit", "--amend").Run()
   x.Check(e)
}
