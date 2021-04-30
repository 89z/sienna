package main

import (
   "github.com/89z/rosso"
   "github.com/89z/youtube"
   "net/url"
   "os"
   "time"
)

func main() {
   if len(os.Args) != 2 {
      println("youtube-views <URL>")
      os.Exit(1)
   }
   arg := os.Args[1]
   addr, err := url.Parse(arg)
   if err != nil {
      panic(err)
   }
   id := addr.Query().Get("v")
   video, err := youtube.NewVideo(id)
   if err != nil {
      panic(err)
   }
   date, err := video.PublishDate()
   if err != nil {
      panic(err)
   }
   view /= time.Since(date).Hours() / 24 / 365
   if view > 10_000_000 {
      rosso.LogFail("Fail", view)
   } else {
      rosso.LogPass("Pass", view)
   }
}
