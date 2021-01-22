package main

import (
   "bufio"
   "bytes"
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

func output(command ...string) (*bufio.Scanner, error) {
   name, arg := command[0], command[1:]
   b, e := exec.Command(name, arg...).Output()
   return bufio.NewScanner(bytes.NewReader(b)), e
}

func system(command ...string) error {
   name, arg := command[0], command[1:]
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   return c.Run()
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
   e := system("go", "mod", "init", "deps")
   check(e)
   e = system("go", "get", mod)
   check(e)
   dep, e := output("go", "list", "-deps", mod + "/...")
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
