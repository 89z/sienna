package main

import (
   "fmt"
   "github.com/89z/rosso"
   "github.com/89z/rosso/musicbrainz"
   "github.com/89z/youtube"
   "io"
   "net/http"
   "net/url"
   "time"
)

func youtubeResult(query string) (string, error) {
   req, err := http.NewRequest("GET", "http://youtube.com/results", nil)
   if err != nil { return "", err }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   rosso.LogInfo("GET", req.URL)
   res, err := new(http.Client).Do(req)
   if err != nil { return "", err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return "", err }
   find, err := rosso.FindSubmatch("/vi/([^/]*)/", body)
   if err != nil { return "", err }
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
      rosso.LogFail("Fail", numberFormat(perYear))
   } else {
      rosso.LogPass("Pass", numberFormat(perYear))
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
