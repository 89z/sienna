package main

import (
   "github.com/89z/rosso"
   "io"
   "net/http"
   "os"
)

func open(source string) (string, error) {
   rosso.LogInfo("Get", source)
   get, err := http.Get(source)
   if err != nil { return "", err }
   body, err := io.ReadAll(get.Body)
   if err != nil { return "", err }
   og, err := rosso.FindSubmatch(`="og:image" content="([^"]+)"`, body)
   if err != nil { return "", err }
   return string(og[1]), nil
}

func main() {
   if len(os.Args) != 2 {
      println("open-graph <URL>")
      os.Exit(1)
   }
   arg := os.Args[1]
   image, err := open(arg)
   if err != nil {
      panic(err)
   }
   println(image)
}
