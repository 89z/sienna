package main

import (
   "fmt"
   "github.com/89z/rosso/musicbrainz"
   "github.com/89z/youtube"
   "io"
   "net/http"
   "net/url"
   "os"
   "regexp"
   "strings"
   "time"
)

const (
   reset = "\x1b[m"
   green = "\x1b[30;102m"
   invert = "\x1b[7m"
   red = "\x1b[30;101m"
)

func youtubeResult(query string) (string, error) {
   req, err := http.NewRequest("GET", "http://youtube.com/results", nil)
   if err != nil { return "", err }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Client).Do(req)
   if err != nil { return "", err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return "", err }
   re := regexp.MustCompile("/vi/([^/]*)/")
   find := re.FindSubmatch(body)
   if find == nil {
      return "", fmt.Errorf("FindSubmatch %v", re)
   }
   return string(find[1]), nil
}

func viewMusicbrainz(addr string) error {
   album, err := musicbrainz.NewRelease(addr)
   if err != nil { return err }
   var artists string
   for _, artist := range album.ArtistCredit { artists += artist.Name + " " }
   for _, media := range album.Media {
      for _, track := range media.Tracks {
         id, err := youtubeResult(artists + track.Title)
         if err != nil { return err }
         vid, err := youtube.NewVideo(id)
         if err != nil { return err }
         err = sinceHours(vid.ViewCount(), vid.PublishDate())
         if err != nil { return err }
         time.Sleep(100 * time.Millisecond)
      }
   }
   return nil
}

func numberFormat(d float64) string {
   var e int
   for d >= 1000 {
      d /= 1000
      e++
   }
   return fmt.Sprintf("%.3f", d) + []string{"", " k", " M", " G"}[e]
}

func sinceHours(view int, date string) error {
   d, err := time.Parse(time.RFC3339[:10], date)
   if err != nil { return err }
   perYear := float64(view) * 24 * 365 / time.Since(d).Hours()
   if perYear > 10_000_000 {
      fmt.Println(red, "fail", reset, numberFormat(perYear))
   } else {
      fmt.Println(green, "pass", reset, numberFormat(perYear))
   }
   return nil
}

func viewYouTube(addr string) error {
   p, err := url.Parse(addr)
   if err != nil { return err }
   id := p.Query().Get("v")
   vid, err := youtube.NewVideo(id)
   if err != nil { return err }
   return sinceHours(vid.ViewCount(), vid.PublishDate())
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println(`usage:
youtube-views <URL>

examples:
https://www.youtube.com/watch?v=6e5cNaU1h1I
https://musicbrainz.org/release/7a629d52-6a61-3ea1-a0a0-dd50bdef63b4
https://musicbrainz.org/release-group/d03bb6b1-d7b4-38ea-974e-847cbb31dca4`)
      os.Exit(1)
   }
   arg := os.Args[1]
   switch {
   case strings.Contains(arg, "musicbrainz.org/"):
      viewMusicbrainz(arg)
   case strings.Contains(arg, "youtube.com/"):
      viewYouTube(arg)
   }
}
