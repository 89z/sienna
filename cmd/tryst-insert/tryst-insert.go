package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "io"
   "net/http"
   "os"
   "regexp"
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
   fmt.Println(invert, "Get", reset, addr)
   res, err := http.Get(addr)
   if err != nil {
      return tableRow{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return tableRow{}, err
   }
   re := regexp.MustCompile("<h1>([^<]+)</h1>")
   title := re.FindSubmatch(body)
   if title == nil {
      return tableRow{}, fmt.Errorf("FindSubmatch %v", re)
   }
   return tableRow{
      A: addr, Date: time.Now(), Title: string(title[1]),
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
