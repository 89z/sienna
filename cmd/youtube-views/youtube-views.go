package main

import (
   "fmt"
   "github.com/89z/musicbrainz"
   "github.com/89z/youtube"
   "io"
   "net/http"
   "net/url"
   "os"
   "path"
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

func numberFormat(d float64) string {
   var e int
   for d >= 1000 {
      d /= 1000
      e++
   }
   return fmt.Sprintf("%.3f", d) + []string{1: " k", " M", " G"}[e]
}

func sinceHours(view int, date string) error {
   d, err := time.Parse(time.RFC3339[:10], date)
   if err != nil { return err }
   perYear := float64(view) * 24 * 365 / time.Since(d).Hours()
   if perYear > 10_000_000 {
      fmt.Println(red, "Fail", reset, numberFormat(perYear))
   } else {
      fmt.Println(green, "Pass", reset, numberFormat(perYear))
   }
   return nil
}

func viewMusicbrainz(r musicbrainz.Release) error {
   var artists string
   for _, artist := range r.ArtistCredit {
      artists += artist.Name + " "
   }
   for _, media := range r.Media {
      for _, track := range media.Tracks {
         id, err := youtubeResult(artists + track.Title)
         if err != nil { return err }
         vid, err := youtube.NewVideo(id)
         if err != nil { return err }
         if err := sinceHours(vid.ViewCount(), vid.PublishDate()); err != nil {
            return err
         }
         time.Sleep(100 * time.Millisecond)
      }
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

func youtubeResult(query string) (string, error) {
   req, err := http.NewRequest("GET", "https://www.youtube.com/results", nil)
   if err != nil { return "", err }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil { return "", err }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return "", fmt.Errorf("Status %v", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil { return "", err }
   re := regexp.MustCompile("/vi/([^/]+)/")
   find := re.FindSubmatch(body)
   if find == nil {
      return "", fmt.Errorf("FindSubmatch %v", re)
   }
   return string(find[1]), nil
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println(`youtube-views <URL>

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
