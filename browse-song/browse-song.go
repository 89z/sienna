package main

import (
   "fmt"
   "log"
   "net/url"
   "os"
   "os/exec"
)

func main() {
   if len(os.Args) != 3 {
      println("browse-song <artist> <song>")
      os.Exit(1)
   }

   band_s := os.Args[1]
   song_s := os.Args[2]
   browse_s := os.Getenv("BROWSER")
   query_s := fmt.Sprintf(`intitle:"%v" intext:"%v topic"`, song_s, band_s)
   m := url.Values{}
   m.Set("q", query_s)
   url_s := "youtube.com/results?" + m.Encode()

   e := exec.Command(browse_s, url_s).Start()
   if e != nil {
      log.Fatal(e)
   }
}
