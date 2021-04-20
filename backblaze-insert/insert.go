package main

import (
   "encoding/json"
   "log"
   "net/url"
   "os"
   "strconv"
   "time"
)

type tableRow struct { Q, S string }

func newTableRow(year, file string) tableRow {
   val := make(url.Values)
   unix := time.Now().Unix()
   // Bryan Adams - Have You Ever Really Loved A Woman.mp3
   // audio id
   format := strconv.FormatInt(unix + 1, 36)
   val.Set("a", format)
   println(format + ".mp4a")
   // image id
   format = strconv.FormatInt(unix, 36)
   val.Set("b", format)
   println(format + ".jpg")
   // platform
   val.Set("p", "mp4a")
   // year
   val.Set("y", year)
   // return
   return tableRow{
      val.Encode(), title,
   }
}

func main() {
   if len(os.Args) != 3 {
      println("backblaze-insert <year> <file>")
      os.Exit(1)
   }
   row := newTableRow(os.Args[1], os.Args[2])
   umber := os.Getenv("UMBER")
   // save
   file, e := os.Open(umber)
   if e != nil {
      log.Fatal(e)
   }
   var rows []tableRow
   json.NewDecoder(file).Decode(&rows)
   // append
   rows = append([]tableRow{row}, rows...)
   // encode
   file, e = os.Create(umber)
   if e != nil {
      log.Fatal(e)
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(rows)
}
