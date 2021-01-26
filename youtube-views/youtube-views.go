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
   url_s := os.Args[1]
   o, e := url.Parse(url_s)
   if e != nil {
      log.Fatal(e)
   }
   id := o.Query().Get("v")
   info, e := youtube.Info(id)
   if e != nil {
      log.Fatal(e)
   }
   view, e := youtube.Views(info)
   if e != nil {
      log.Fatal(e)
   }
   color, _ := youtube.Color(view)
   println(color)
}
