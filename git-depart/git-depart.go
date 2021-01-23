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

func green(s string) string {
   return "\x1b[92m" + s + "\x1b[m"
}

func main() {
   e := x.System("git", "commit", "--verbose")
   check(e)
   if x.IsFile("config.toml") {
      println(green("remove docs"))
      os.RemoveAll("docs")
      println(green("hugo"))
      e = x.System("hugo")
      check(e)
      println(green("git add"))
      e = x.System("git", "add", ".")
      check(e)
      println(green("git commit"))
      e = x.System("git", "commit", "--amend")
      check(e)
   }
   println(green("git push"))
}
