package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "net/http"
   "net/url"
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
