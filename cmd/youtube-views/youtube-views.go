package main

import (
   "github.com/89z/x/youtube"
   "net/url"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("youtube-views <URL>")
      os.Exit(1)
   }
   arg := os.Args[1]
   addr, e := url.Parse(arg)
   if e != nil {
      panic(e)
   }
   id := addr.Query().Get("v")
   info, e := youtube.Info(id)
   if e != nil {
      panic(e)
   }
   info.Views()
}
