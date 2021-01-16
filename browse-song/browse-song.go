//go:generate mkwinsyscall -output zbrowse-song.go browse-song.go
//sys shellExecute(hwnd int, verb string, file string, args string, cwd string, showCmd int) (err error) = shell32.ShellExecuteW
package main

import (
   "fmt"
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
   queryVal := url.Values{}
   queryVal.Set("q", query)
   result := "https://www.youtube.com/results?" + queryVal.Encode()
   e := shellExecute(0, "", result, "", "", sw_shownormal)
   if e != nil {
      log.Fatal(e)
   }
}
