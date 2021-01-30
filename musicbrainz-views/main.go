package main

import (
   "github.com/89z/x"
   "github.com/89z/x/musicbrainz"
   "github.com/89z/x/youtube"
   "os"
   "strings"
   "time"
)

func main() {
   if len(os.Args) != 2 {
      println(`usage:
musicbrainz-views <URL>

examples:
https://musicbrainz.org/release-group/d03bb6b1-d7b4-38ea-974e-847cbb31dca4
https://musicbrainz.org/release/7a629d52-6a61-3ea1-a0a0-dd50bdef63b4`)
      os.Exit(1)
   }
   album, e := musicbrainz.NewRelease(os.Args[1])
   x.Check(e)
   var out []string
   for _, artist := range album.ArtistCredit {
      out = append(out, artist.Name)
   }
   artist := strings.Join(out, " ")
   for _, media := range album.Media {
      for _, track := range media.Tracks {
         id, e := youtubeResult(artist + " " + track.Title)
         x.Check(e)
         info, e := youtube.Info(string(id))
         x.Check(e)
         views, e := info.Views()
         x.Check(e)
         println(views)
         time.Sleep(100 * time.Millisecond)
      }
   }
}
