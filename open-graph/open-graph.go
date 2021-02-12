package main

import (
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "net/http"
   "os"
)

func open(source string) (string, error) {
   println(x.ColorCyan("Get"), source)
   get, e := http.Get(source)
   if e != nil {
      return "", e
   }
   body, e := ioutil.ReadAll(get.Body)
   if e != nil {
      return "", e
   }
   return string(
      x.FindSubmatch(`="og:image" content="([^"]+)"`, body),
   ), nil
}

func main() {
   if len(os.Args) != 2 {
      println("open-graph <URL>")
      os.Exit(1)
   }
   image, e := open(os.Args[1])
   if e != nil {
      log.Fatal(e)
   }
   println(image)
}
