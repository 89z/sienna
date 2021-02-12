package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "net/http"
   "os"
)

func open(url string) (string, error) {
   get, e := http.Get(url)
   if e != nil {
      return "", e
   }
   body, e := ioutil.ReadAll(get.Body)
   if e != nil {
      return "", e
   }
   image := x.FindSubmatch(`="og:image" content="([^"]+)"`, body)
   return string(image), nil
}

func main() {
   if len(os.Args) != 2 {
      println("open-graph <URL>")
      os.Exit(1)
   }
   url, e := open(os.Args[1])
   if e != nil {
      log.Fatal(e)
   }
   println(url)
}
