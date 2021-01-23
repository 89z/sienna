package main

import (
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

func isFile(name string) bool {
   fi, e := os.Stat(name)
   return e == nil && fi.Mode().IsRegular()
}

func system(name string, arg ...string) error {
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   return c.Run()
}

func main() {
   e := system("git", "commit", "--verbose")
   check(e)
   if isFile("config.toml") {
      println(green("remove docs"))
      os.RemoveAll("docs")
      println(green("hugo"))
      e = system("hugo")
      check(e)
      println(green("git add"))
      e = system("git", "add", ".")
      check(e)
      println(green("git commit"))
      e = system("git", "commit", "--amend")
      check(e)
   }
   println(green("git push"))
}
