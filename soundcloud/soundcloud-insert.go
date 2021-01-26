package main

import (
   "github.com/89z/x/soundcloud"
   "log"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("soundcloud-insert <URL>")
      os.Exit(1)
   }
   rec, e := soundcloud.Insert(os.Args[1])
   if e != nil {
      log.Fatal(e)
   }
   println(rec)
}
