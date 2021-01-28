package main

import (
   "github.com/89z/x"
   "net/url"
)

func youtubeResult(query string) ([]byte, error) {
   value := make(url.Values)
   value.Set("search_query", query)
   body, e := x.GetContents(
      "https://www.youtube.com/results?" + value.Encode(),
   )
   if e != nil {
      return nil, e
   }
   return x.FindSubmatch("/vi/([^/]*)/", body), nil
}
