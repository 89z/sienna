package main

import (
   "fmt"
   "log"
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
   url_s := fmt.Sprintf(
      `youtube.com/results?q=intext:"%v - topic" intitle:"%v"`, band_s, song_s,
   )

   e := exec.Command(browse_s, url_s).Start()
   if e != nil {
      log.Fatal(e)
   }
}
