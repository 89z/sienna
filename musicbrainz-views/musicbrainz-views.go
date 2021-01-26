package main

import (
   "github.com/89z/x"
   "github.com/89z/x/musicbrainz"
   "github.com/89z/x/youtube"
   "io/ioutil"
   "net/http"
   "net/url"
   "os"
   "strings"
   "time"
)

func youtubeResult(query string) (string, error) {
   value := url.Values{}
   value.Set("search_query", query)
   result := "https://www.youtube.com/results?" + value.Encode()
   resp, e := http.Get(result)
   if e != nil {
      return "", e
   }
   body, e := ioutil.ReadAll(resp.Body)
   if e != nil {
      return "", e
   }
   return x.FindSubmatch("/vi/([^/]*)/", body), nil
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
   url := os.Args[1]
   album, e := musicbrainz.Release(url)
   x.Check(e)
   out_a := []string{}
   artist_a := album.A("artist-credit")
   for n := range artist_a {
      artist_s := artist_a.M(n).S("name")
      out_a = append(out_a, artist_s)
   }
   artist_s := strings.Join(out_a, " ")
   media_a := album.A("media")
   for n := range media_a {
      track_a := media_a.M(n).A("tracks")
      for n := range track_a {
         title_s := track_a.M(n).S("title")
         ytid, e := youtubeResult(artist_s + " " + title_s)
         x.Check(e)
         info_m, e := youtube.Info(ytid)
         x.Check(e)
         view_n, e := youtube.Views(info_m)
         x.Check(e)
         color_s, b := youtube.Color(view_n)
         println(color_s)
         if b {
            print("youtube.com/watch?v=", ytid, "\n")
         }
         time.Sleep(500 * time.Millisecond)
      }
   }
}
