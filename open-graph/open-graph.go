package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "regexp"
)

var patterns = []string{
   ` content="([^"]+)" property="og:image"`,
   `="og:image" content="([^"]+)"`,
   `="og:video" content="([^"]+)"`,
   `=og:image content="([^"]+)"`,
}

func findSubmatch(re string, input []byte) string {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return string(a[1])
}

func open(url string) (urls []string, e error) {
   get, e := http.Get(url)
   if e != nil {
      return nil, e
   }
   body, e := ioutil.ReadAll(get.Body)
   if e != nil {
      return nil, e
   }
   for _, re := range patterns {
      sub := findSubmatch(re, body)
      if sub != "" {
         urls = append(urls, sub)
      }
   }
   return
}

func main() {
   if len(os.Args) != 2 {
      println("open-graph <URL>")
      os.Exit(1)
   }
   url := os.Args[1]
   urls, e := open(url)
   if e != nil {
      log.Fatal(e)
   }
   for _, url := range urls {
      println(url)
   }
}
