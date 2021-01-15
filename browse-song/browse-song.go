//go:generate mkwinsyscall -output zbrowse-song.go browse-song.go
//sys ShellExecute(hwnd int, verb string, file string, args string, cwd string, showCmd int) (err error) = shell32.ShellExecuteW
package main

import (
   "fmt"
   "log"
   "net/url"
   "os"
)

const SW_SHOWNORMAL = 1

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
   e := ShellExecute(0, "", result, "", "", SW_SHOWNORMAL)
   if e != nil {
      log.Fatal(e)
   }
}
