package main

import (
   "bufio"
   "bytes"
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
   "strings"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func main() {
   if len(os.Args) != 2 {
      println(`usage:
go-deps <URL>

example:
   go-deps https://github.com/dinedal/textql`)
      os.Exit(1)
   }
   mod := os.Args[1][8:]
   e := x.System("go", "mod", "init", "deps")
   check(e)
   e = x.System("go", "get", mod)
   check(e)
   dep, e := x.Popen("go", "list", "-deps", mod + "/...")
   check(e)
   os.Remove("go.mod")
   os.Remove("go.sum")
   deps := 0
   for dep.Scan() {
      text := dep.Text()
      if strings.Contains(text, "/internal/") {
         continue
      }
      if ! strings.Contains(text, ".") {
         continue
      }
      if strings.HasPrefix(text, mod + "/") {
         continue
      }
      println(text)
      deps++
   }
   print("\n", deps, " deps\n")
}
