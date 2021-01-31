package main

import (
   "fmt"
   "github.com/89z/x"
   "github.com/89z/x/sys"
   "os"
)

const sw_shownormal = 1

func main() {
   if len(os.Args) != 3 {
      println("browse-song <artist> <song>")
      os.Exit(1)
   }
   artist, song := os.Args[1], os.Args[2]
   url := x.NewURL()
   url.Host = "youtube.com"
   url.Path = "results"
   url.QuerySet(
      "q", fmt.Sprintf(`intext:"%v topic" intitle:"%v"`, artist, song),
   )
   e := sys.ShellExecute(
      0, "", url.String(), "", "", sw_shownormal,
   )
   x.Check(e)
}
