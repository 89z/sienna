package main

import (
   "io/ioutil"
   "net/http"
   "net/url"
   "regexp"
)

func findSubmatch(re string, input []byte) string {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return string(a[1])
}

func youtubeResult(query string) (string, error) {
   value := url.Values{}
   value.Set("search_query", query)
   result := "https://www.youtube.com/results?" + value.Encode()
   resp, e := http.Get(result)
   if e != nil {
      return "", e
   }
   get, e := ioutil.ReadAll(resp.Body)
   if e != nil {
      return "", e
   }
   return findSubmatch("/vi/([^/]*)/", get), nil
}
