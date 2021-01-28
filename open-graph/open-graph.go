package main

import (
   "github.com/89z/x"
   "os"
)

func open(url string) (string, error) {
   get, e := x.GetContents(url)
   if e != nil {
      return "", e
   }
   return x.FindSubmatch(`="og:image" content="([^"]+)"`, get), nil
}

func main() {
   if len(os.Args) != 2 {
      println("open-graph <URL>")
      os.Exit(1)
   }
   url, e := open(os.Args[1])
   x.Check(e)
   println(url)
}
