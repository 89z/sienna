package main

import (
   "github.com/89z/x/youtube"
   "log"
   "net/url"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("youtube-views <URL>")
      os.Exit(1)
   }
   source, e := url.Parse(os.Args[1])
   if e != nil {
      log.Fatal(e)
   }
   id := source.Query().Get("v")
   info, e := youtube.Info(id)
   if e != nil {
      log.Fatal(e)
   }
   views, e := info.Views()
   if e != nil {
      log.Fatal(e)
   }
   println(views)
}
