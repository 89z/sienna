package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/pelletier/go-toml"
   "net/http"
   "net/url"
   "os"
   "time"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

type tableRow struct {
   A string
   Categories []string
   Date time.Time
   Img string
   Tags []string
   Title string
}

func newTableRow(addr string) (tableRow, error) {
   // A
   parse, err := url.Parse(addr)
   if err != nil {
      return tableRow{}, err
   }
   parse.Fragment = ""
   // Title
   fmt.Println(invert, "Get", reset, addr)
   res, err := http.Get(addr)
   if err != nil {
      return tableRow{}, err
   }
   defer res.Body.Close()
   doc, err := mech.NewNode(res.Body)
   if err != nil {
      return tableRow{}, err
   }
   // return
   return tableRow{
      A: parse.String(), Date: time.Now(), Title: doc.ByTag("h1").Text(),
   }, nil
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("tryst-insert <URL>")
      return
   }
   addr := os.Args[1]
   row, err := newTableRow(addr)
   if err != nil {
      panic(err)
   }
   file, err := os.Create("index.md")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   file.WriteString("+++\n")
   toml.NewEncoder(file).Encode(row)
   file.WriteString("+++\n")
}
