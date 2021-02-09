package main

import (
   "github.com/89z/x"
   "os"
   "os/exec"
)

func main() {
   e := exec.Command("git", "commit", "--verbose").Run()
   x.Check(e)
   _, e = os.Stat("config.toml")
   if e != nil {
      return
   }
   println(x.ColorGreen("remove docs"))
   os.RemoveAll("docs")
   println(x.ColorGreen("hugo"))
   e = exec.Command("hugo").Run()
   x.Check(e)
   println(x.ColorGreen("git add"))
   e = exec.Command("git", "add", ".").Run()
   x.Check(e)
   println(x.ColorGreen("git commit"))
   e = exec.Command("git", "commit", "--amend").Run()
   x.Check(e)
}
