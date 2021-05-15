package main

import (
   "encoding/json"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

func split(s string) (string, string) {
   f := strings.LastIndexByte
   d, e := f(s, os.PathSeparator), f(s, '.')
   return s[d + 1:e], s[e + 1:]
}

type tableRow struct { Q, S string }

func newTableRow(year, file string) tableRow {
   stem, ext := split(file)
   val := make(url.Values)
   unix := time.Now().Unix()
   // audio id
   format := strconv.FormatInt(unix + 1, 36)
   val.Set("a", format)
   println(format + "." + ext)
   // image id
   format = strconv.FormatInt(unix, 36)
   val.Set("b", format)
   println(format + ".jpg")
   // platform
   val.Set("p", ext)
   // year
   val.Set("y", year)
   // return
   return tableRow{
      val.Encode(), stem,
   }
}

func main() {
   if len(os.Args) != 3 {
      println("backblaze-insert <year> <file>")
      return
   }
   row := newTableRow(os.Args[1], os.Args[2])
   umber := os.Getenv("UMBER")
   // save
   file, e := os.Open(umber)
   if e != nil {
      panic(e)
   }
   var rows []tableRow
   json.NewDecoder(file).Decode(&rows)
   // append
   rows = append([]tableRow{row}, rows...)
   // encode
   file, e = os.Create(umber)
   if e != nil {
      panic(e)
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(rows)
}
