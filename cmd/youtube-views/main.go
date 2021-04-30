package main

import (
   "os"
   "strings"
)

func main() {
   if len(os.Args) != 2 {
      println(`usage:
youtube-views <URL>

examples:
https://www.youtube.com/watch?v=6e5cNaU1h1I
https://musicbrainz.org/release/7a629d52-6a61-3ea1-a0a0-dd50bdef63b4
https://musicbrainz.org/release-group/d03bb6b1-d7b4-38ea-974e-847cbb31dca4`)
      os.Exit(1)
   }
   arg := os.Args[1]
   switch {
   case strings.HasPrefix(arg, "https://musicbrainz.org/"):
      viewMusicbrainz(arg)
   case strings.HasPrefix(arg, "https://www.youtube.com/"):
      viewYouTube(arg)
   }
}
