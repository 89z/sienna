package main

import (
   "fmt"
   "io/ioutil"
   "log"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      fmt.Println("code-points <file>")
      os.Exit(1)
   }
   s := os.Args[1]
   y, e := ioutil.ReadFile(s)
   if e != nil {
      log.Fatal(e)
   }
   for _, r := range string(y) {
      if r == '\n' {
         fmt.Println("--------------------------------------------------------")
         continue
      }
      fmt.Printf("%c\tU+%04X\n", r, r)
   }
}
