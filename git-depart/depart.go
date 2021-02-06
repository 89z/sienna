package main

import (
   "github.com/89z/x"
   "os"
)

func main() {
   e := x.Command("git", "commit", "--verbose").Run()
   x.Check(e)
   if x.IsFile("config.toml") {
      println(x.ColorGreen("remove docs"))
      os.RemoveAll("docs")
      println(x.ColorGreen("hugo"))
      e = x.Command("hugo").Run()
      x.Check(e)
      println(x.ColorGreen("git add"))
      e = x.Command("git", "add", ".").Run()
      x.Check(e)
      println(x.ColorGreen("git commit"))
      e = x.Command("git", "commit", "--amend").Run()
      x.Check(e)
   }
}
