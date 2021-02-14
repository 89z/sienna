package main

import (
   "fmt"
   "github.com/89z/x/sys"
   "log"
   "net/url"
   "os"
)

const sw_shownormal = 1

func main() {
   if len(os.Args) != 3 {
      println("browse-song <artist> <song>")
      os.Exit(1)
   }
   artist, song := os.Args[1], os.Args[2]
   query := fmt.Sprintf(`intext:"%v topic" intitle:"%v"`, artist, song)
   value := make(url.Values)
   value.Set("q", query)
   e := sys.ShellExecute(
      0,
      "",
      "http://youtube.com/results?" + value.Encode(),
      "",
      "",
      sw_shownormal,
   )
   if e != nil {
      log.Fatal(e)
   }
}
