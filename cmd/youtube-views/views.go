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

func youtubeResult(query string) (string, error) {
   req, e := http.NewRequest("GET", "http://youtube.com/results", nil)
   if e != nil { return "", e }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   rosso.LogInfo("GET", req.URL)
   res, e := new(http.Client).Do(req)
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
