package main

import (
   "log"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("go-repos <count | download>")
      os.Exit(1)
   }
   var e error
   if os.Args[1] == "download" {
      e = download()
   } else {
      e = count()
   }
   if e != nil {
      log.Fatal(e)
   }
}
