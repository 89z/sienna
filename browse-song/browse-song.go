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
   bandArg := os.Args[1]
   songArg := os.Args[2]
   query := fmt.Sprintf(`intitle:"%v" intext:"%v topic"`, songArg, bandArg)
   value := url.Values{}
   value.Set("q", query)
   result := "https://www.youtube.com/results?" + value.Encode()
   e := sys.ShellExecute(0, "", result, "", "", sw_shownormal)
   if e != nil {
      log.Fatal(e)
   }
}
