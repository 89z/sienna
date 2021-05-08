package main

import (
   "fmt"
   "io"
   "net/http"
   "os"
   "regexp"
)

func open(source string) (string, error) {
   fmt.Println("\x1b[7m GET \x1b[m", source)
   get, err := http.Get(source)
   if err != nil { return "", err }
   body, err := io.ReadAll(get.Body)
   if err != nil { return "", err }
   re := regexp.MustCompile(`="og:image" content="([^"]+)"`)
   image := re.FindSubmatch(body)
   if image == nil {
      return "", fmt.Errorf("FindSubmatch %v", re)
   }
   return string(image[1]), nil
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("open-graph <URL>")
      os.Exit(1)
   }
   arg := os.Args[1]
   image, err := open(arg)
   if err != nil {
      panic(err)
   }
   fmt.Println(image)
}
