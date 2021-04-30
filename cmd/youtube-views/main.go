package main

import (
   "bytes"
   "fmt"
   "github.com/89z/rosso"
   "github.com/89z/rosso/musicbrainz"
   "github.com/89z/youtube"
   "net/http"
   "net/url"
   "os"
   "regexp"
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
   view := time.Since(date).Hours() / 24 / 365
   if view > 10_000_000 {
      rosso.LogFail("Fail", view)
   } else {
      rosso.LogPass("Pass", view)
   }
}

func main() {
   if len(os.Args) != 2 {
      println(`usage:
musicbrainz-views <URL>

examples:
https://musicbrainz.org/release-group/d03bb6b1-d7b4-38ea-974e-847cbb31dca4
https://musicbrainz.org/release/7a629d52-6a61-3ea1-a0a0-dd50bdef63b4`)
      os.Exit(1)
   }
   arg := os.Args[1]
   album, e := musicbrainz.NewRelease(arg)
   if e != nil {
      panic(e)
   }
   var artists string
   for _, artist := range album.ArtistCredit { artists += artist.Name + " " }
   for _, media := range album.Media {
      for _, track := range media.Tracks {
         id, e := youtubeResult(artists + track.Title)
         if e != nil {
            panic(e)
         }
         info, e := youtube.Info(id)
         if e != nil {
            panic(e)
         }
         info.Views()
         time.Sleep(100 * time.Millisecond)
      }
   }
}
