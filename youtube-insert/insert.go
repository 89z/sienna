package main

import (
   "encoding/json"
   "github.com/89z/x"
   "github.com/89z/x/youtube"
   "net/http"
   "net/url"
   "os"
   "regexp"
   "strconv"
   "strings"
   "time"
)

/* the order doesnt matter here, as we will find the lowest date of all
matches */
var patterns = []string{
   ` (\d{4})`, `(\d{4}) `, `Released on: (\d{4})`, `℗ (\d{4})`,
}

func findSubmatch(re, input string) string {
   a := regexp.MustCompile(re).FindStringSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}

func getImage(id string) string {
   url := "https://i.ytimg.com/vi/"
   if httpHead(url + id + "/sddefault.jpg") {
      return ""
   }
   if httpHead(url + id + "/sd1.jpg") {
      return "sd1"
   }
   return "hqdefault"
}

func httpHead(url string) bool {
   println(url)
   resp, e := http.Head(url)
   return e == nil && resp.StatusCode == 200
}

func marshal(v interface{}) ([]byte, error) {
   var dst bytes.Buffer
   enc := json.NewEncoder(&dst)
   enc.SetEscapeHTML(false)
   err := enc.Encode(v)
   if err != nil {
      return nil, err
   }
   return dst.Bytes()[:dst.Len() - 1], nil
}
