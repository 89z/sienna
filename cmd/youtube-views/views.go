package main

import (
   "bytes"
   "fmt"
   "github.com/89z/rosso"
   "github.com/89z/rosso/musicbrainz"
   "github.com/89z/youtube"
   "net/http"
   "net/url"
   "regexp"
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
   var buf bytes.Buffer
   buf.ReadFrom(res.Body)
   str := buf.String()
   find := regexp.MustCompile("/vi/([^/]*)/").FindStringSubmatch(str)
   if len(find) < 2 {
      return "", fmt.Errorf("%v", req.URL)
   }
   return find[1], nil
}

func viewMusicbrainz(addr string) error {
   album, err := musicbrainz.NewRelease(arg)
   if err != nil { return err }
   var artists string
   for _, artist := range album.ArtistCredit { artists += artist.Name + " " }
   for _, media := range album.Media {
      for _, track := range media.Tracks {
         id, err := youtubeResult(artists + track.Title)
         if err != nil { return err }
         info, err := youtube.Info(id)
         if err != nil { return err }
         info.Views()
         time.Sleep(100 * time.Millisecond)
      }
   }
   return nil
}

func viewYouTube(addr string) error {
   addr, err := url.Parse(arg)
   if err != nil { return err }
   id := addr.Query().Get("v")
   video, err := youtube.NewVideo(id)
   if err != nil { return err }
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
