package main

import (
   "github.com/89z/x"
   "github.com/89z/x/youtube"
   "io/ioutil"
   "net/http"
)

func infoViews(id string) (string, error) {
   info, e := youtube.Info(id)
   if e != nil {
      return "", e
   }
   return info.Views()
}

func youtubeResult(query string) (string, error) {
   url := x.NewURL()
   url.Host = "youtube.com"
   url.Path = "results"
   url.Query.Set("search_query", query)
   get, e := http.Get(url.String())
   if e != nil {
      return "", e
   }
   body, e := ioutil.ReadAll(get.Body)
   if e != nil {
      return "", e
   }
   return string(
      x.FindSubmatch("/vi/([^/]*)/", body),
   ), nil
}
