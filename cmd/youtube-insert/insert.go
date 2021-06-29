package main

import (
   "encoding/json"
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
   "math/big"
   "net/url"
   "os"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func mod(x, y int64) int64 {
   bx, by := big.NewInt(x), big.NewInt(y)
   return new(big.Int).Mod(bx, by).Int64()
}

func score(i youtube.Image) int64 {
   return mod(480 - i.Height, 720) + i.Frame + i.Format
}

type tableRow struct {
   Q string
   S string
}

func main() {
   var dry bool
   flag.BoolVar(&dry, "d", false, "dry run")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("youtube-insert [flag] <URL>")
      flag.PrintDefaults()
      return
   }
   addr, err := url.Parse(flag.Arg(0))
   if err != nil {
      panic(err)
   }
   row, err := newTableRow(addr.Query().Get("v"))
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(row)
   if dry {
      return
   }
   // write
   umber := os.Getenv("UMBER")
   // decode
   var rows []*tableRow
   file, err := os.Open(umber)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   json.NewDecoder(file).Decode(&rows)
   // append
   rows = append([]*tableRow{row}, rows...)
   // encode
   if file, err := os.Create(umber); err != nil {
      panic(err)
   } else {
      defer file.Close()
      enc := json.NewEncoder(file)
      enc.SetEscapeHTML(false)
      enc.SetIndent("", " ")
      enc.Encode(rows)
   }
}
