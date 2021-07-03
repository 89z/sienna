package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "os"
)

type newsArticle struct {
   Video struct {
      ContentURL string
   }
}

func open(source string) ([]string, error) {
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
   script := doc.ByAttr("type", "application/ld+json")
   for script.Scan() {
      text := []byte(script.Text())
      var na newsArticle
      json.Unmarshal(text, &na)
      if na.Video.ContentURL != "" {
         nodes = append(nodes, na.Video.ContentURL)
      }
   }
   return nodes, nil
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("page-media <URL>")
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
