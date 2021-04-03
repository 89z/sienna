package main

import (
   "bytes"
   "fmt"
   "github.com/89z/x/musicbrainz"
   "github.com/89z/x/youtube"
   "log"
   "net/http"
   "os"
   "regexp"
   "time"
)

func youtubeResult(query string) (string, error) {
   req, e := http.NewRequest("GET", "http://youtube.com/results", nil)
   if e != nil { return "", e }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   res, e := http.DefaultClient.Do(req)
   if e != nil { return "", e }
   var buf bytes.Buffer
   buf.ReadFrom(res.Body)
   str := buf.String()
   find := regexp.MustCompile("/vi/([^/]*)/").FindStringSubmatch(str)
   if len(find) < 2 {
      return "", fmt.Errorf("%v", req.URL)
   }
   return find[1], nil
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
   album, e := musicbrainz.NewRelease(os.Args[1])
   if e != nil {
      log.Fatal(e)
   }
   var artist string
   for _, each := range album.ArtistCredit { artist += each.Name + " " }
   for _, media := range album.Media {
      for _, track := range media.Tracks {
         id, e := youtubeResult(artist + track.Title)
         if e != nil {
            log.Fatal(e)
         }
         info, e := youtube.Info(id)
         if e != nil {
            log.Fatal(e)
         }
         info.Views()
         time.Sleep(100 * time.Millisecond)
      }
   }
}
