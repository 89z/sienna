package main

import (
   "bytes"
   "github.com/89z/x"
   "os"
   "strings"
)

func main() {
   files, e := os.ReadDir(".")
   if e != nil {
      panic(e)
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
      read, e := os.ReadFile(name)
      if e != nil {
         panic(e)
      }
      countArg += bytes.Count(read, []byte{'\n'})
   }
   var cmd x.Cmd
   cmd.Run("go", "mod", "init", "init")
   cmd.Run("go", "mod", "tidy")
   sum, e := os.ReadFile("go.sum")
   if e != nil {
      panic(e)
   }
   countSum := bytes.Count(sum, []byte{'\n'})
   println("*.go", countArg)
   println("go.sum", countSum)
   println("product", countSum * countArg)
}
