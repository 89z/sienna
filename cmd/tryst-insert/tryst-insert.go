package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/pelletier/go-toml"
   "net/url"
   "os"
   "time"
)

type tableRow struct {
   A string
   Categories []string
   Date time.Time
   Img string
   Tags []string
   Title string
}

func newTableRow(addr string) (*tableRow, error) {
   // A
   parse, err := url.Parse(addr)
   if err != nil {
      return nil, err
   }
   parse.Fragment = ""
   // Title
   res, err := mech.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   doc, err := mech.Parse(res.Body)
   if err != nil {
      return nil, err
   }
   doc = doc.ByAttr("class", "dts-section-page-heading-title")
   doc.Scan()
   doc = doc.ByTag("h1")
   doc.Scan()
   return &tableRow{
      A: parse.String(), Date: time.Now(), Title: doc.Text(),
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
