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

   query_m := url.Values{"q": []string{
      fmt.Sprintf(`intext:"%v - topic" intitle:"%v"`, band_s, song_s),
   }}
   query_s := query_m.Encode()

   e := exec.Command(browse_s, "youtube.com/results?" + query_s).Start()
   if e != nil {
      log.Fatal(e)
   }
}
