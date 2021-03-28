package main

import (
   "github.com/89z/x"
   "io"
   "log"
   "net/http"
   "os"
   "regexp"
)

func open(source string) (string, error) {
   x.LogInfo("Get", source)
   get, e := http.Get(source)
   if e != nil {
      return "", e
   }
   body, e := io.ReadAll(get.Body)
   if e != nil {
      return "", e
   }
   re := regexp.MustCompile(`="og:image" content="([^"]+)"`)
   return string(re.FindSubmatch(body)[1]), nil
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
