// go-size
package main

import (
   "bytes"
   "os"
   "os/exec"
   "strings"
)

func main() {
   files, err := os.ReadDir(".")
   if err != nil {
      panic(err)
   }
   var countArg int
   for _, file := range files {
      name := file.Name()
      if name == ".git" { continue }
      f := strings.HasSuffix
      if f(name, "_test.go") || ! f(name, ".go") {
         os.RemoveAll(name)
         continue
      }
      println(name)
      read, err := os.ReadFile(name)
      if err != nil {
         panic(err)
      }
      countArg += bytes.Count(read, []byte{'\n'})
   }
   exec.Command("go", "mod", "init", "init").Run()
   exec.Command("go", "mod", "tidy").Run()
   sum, err := os.ReadFile("go.sum")
   if err != nil {
      panic(err)
   }
   countSum := bytes.Count(sum, []byte{'\n'})
   println("*.go", countArg)
   println("go.sum", countSum)
   println("product", countSum * countArg)
}
