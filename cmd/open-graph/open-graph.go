package main

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "os"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func open(source string) ([]string, error) {
   fmt.Println(invert, "Get", reset, source)
   res, err := http.Get(source)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   doc, err := mech.Parse(res.Body)
   if err != nil {
      return nil, err
   }
   var nodes []string
   img := doc.ByAttr("property", "og:image")
   for img.Scan() {
      nodes = append(nodes, img.Attr("content"))
   }
   vid := doc.ByAttr("property", "og:video")
   for vid.Scan() {
      nodes = append(nodes, vid.Attr("content"))
   }
   return nodes, nil
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("open-graph <URL>")
      return
   }
   arg := os.Args[1]
   items, err := open(arg)
   if err != nil {
      panic(err)
   }
   for _, item := range items {
      fmt.Println(item)
   }
}
