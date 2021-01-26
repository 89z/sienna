package main

import (
   "net/http"
   "regexp"
)

func findSubmatch(re, input string) string {
   a := regexp.MustCompile(re).FindStringSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}

func getImage(id_s string) string {
   url_s := "https://i.ytimg.com/vi/"
   if httpHead(url_s + id_s + "/sddefault.jpg") {
      return ""
   }
   if httpHead(url_s + id_s + "/sd1.jpg") {
      return "/sd1"
   }
   return "/hqdefault"
}

func httpHead(url string) bool {
   println(url)
   resp, e := http.Head(url)
   return e == nil && resp.StatusCode == 200
}
