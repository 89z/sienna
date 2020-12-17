package main

import (
   "log"
   "musicdb/youtube"
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
   id_s := o.Query().Get("v")
   m := youtube.Info(id_s)
   view_s := youtube.Views(m)
   println(view_s)
}
