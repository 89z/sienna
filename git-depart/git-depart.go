package main

import (
   "github.com/89z/x"
   "os"
)

func main() {
   e := x.System("git", "commit", "--verbose")
   x.Check(e)
   if x.IsFile("config.toml") {
      println(x.ColorGreen("remove docs"))
      os.RemoveAll("docs")
      println(x.ColorGreen("hugo"))
      e = x.System("hugo")
      x.Check(e)
      println(x.ColorGreen("git add"))
      e = x.System("git", "add", ".")
      x.Check(e)
      println(x.ColorGreen("git commit"))
      e = x.System("git", "commit", "--amend")
      x.Check(e)
   }
   println(x.ColorGreen("git push"))
}
