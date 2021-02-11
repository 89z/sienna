package main

import (
   "github.com/89z/x"
   "net/http"
   "regexp"
)

func getImage(id string) string {
   if httpHead("http://i.ytimg.com/vi/" + id + "/sddefault.jpg") {
      return ""
   }
   if httpHead("http://i.ytimg.com/vi/" + id + "/sd1.jpg") {
      return "sd1"
   }
   return "hqdefault"
}

func httpHead(s string) bool {
   println(x.ColorCyan("Head"), s)
   resp, e := http.Head(s)
   return e == nil && resp.StatusCode == 200
}

func findStringSubmatch(re, input string) string {
   a := regexp.MustCompile(re).FindStringSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}
