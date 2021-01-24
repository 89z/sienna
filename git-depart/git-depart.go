package main

import (
   "github.com/89z/x"
   "log"
   "os"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func main() {
   e := x.System("git", "commit", "--verbose")
   check(e)
   if x.IsFile("config.toml") {
      println(x.ColorGreen("remove docs"))
      os.RemoveAll("docs")
      println(x.ColorGreen("hugo"))
      e = x.System("hugo")
      check(e)
      println(x.ColorGreen("git add"))
      e = x.System("git", "add", ".")
      check(e)
      println(x.ColorGreen("git commit"))
      e = x.System("git", "commit", "--amend")
      check(e)
   }
   println(x.ColorGreen("git push"))
}
