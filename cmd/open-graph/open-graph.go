package main

import (
   "fmt"
   "html"
   "io"
   "net/http"
   "os"
   "regexp"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

var patterns = []string{
   ` content="([^"]+)" property="og:image"`,
   `="og:image" content="([^"]+)"`,
   `="og:video" content="([^"]+)"`,
}

func open(source string) ([]string, error) {
   fmt.Println(invert, "Get", reset, source)
   get, err := http.Get(source)
   if err != nil { return nil, err }
   body, err := io.ReadAll(get.Body)
   if err != nil { return nil, err }
   var results []string
   for _, pattern := range patterns {
      find := regexp.MustCompile(pattern).FindSubmatch(body)
      if find != nil {
         media := string(find[1])
         media = html.UnescapeString(media)
         results = append(results, media)
      }
   }
   return results, nil
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
