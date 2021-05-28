package main

import (
   "fmt"
   "github.com/89z/musicbrainz"
   "os"
   "path"
   "strings"
)

func main() {
   if len(os.Args) != 2 {
      fmt.Println(`usage:
youtube-views <URL>

examples:
https://www.youtube.com/watch?v=6e5cNaU1h1I
https://musicbrainz.org/release/7a629d52-6a61-3ea1-a0a0-dd50bdef63b4
https://musicbrainz.org/release-group/d03bb6b1-d7b4-38ea-974e-847cbb31dca4`)
      return
   }
   arg := os.Args[1]
   if strings.Contains(arg, "youtube.com/") {
      err := viewYouTube(arg)
      if err != nil {
         panic(err)
      }
      return
   }
   id := path.Base(arg)
   if strings.Contains(arg, "musicbrainz.org/release/") {
      r, err := musicbrainz.NewRelease(id)
      if err != nil {
         panic(err)
      }
      if err := viewMusicbrainz(r); err != nil {
         panic(err)
      }
      return
   }
   g, err := musicbrainz.NewGroup(id)
   if err != nil {
      panic(err)
   }
   g.Sort()
   if err := viewMusicbrainz(g.Releases[0]); err != nil {
      panic(err)
   }
}
