package main

import (
   "encoding/json"
   "fmt"
   "net/url"
   "os"
   "path/filepath"
   "strconv"
   "time"
)

type tableRow struct {
   Q string
   S string
}

func newTableRow(year, audio, image string) tableRow {
   val := make(url.Values)
   // year
   val.Set("y", year)
   // audio id
   val.Set("a", strconv.FormatInt(time.Now().Unix(), 36))
   // platform: audio extension
   val.Set("p", filepath.Ext(audio)[1:])
   // image stem
   base := filepath.Base(image)
   val.Set("b", base[:len(base)-4])
   // return
   base = filepath.Base(audio)
   return tableRow{
      val.Encode(),
      // audio stem
      base[:len(base)-4],
   }
}

func main() {
   if len(os.Args) != 4 {
      fmt.Println("backblaze-insert <year> <audio> <image>")
      return
   }
   row := newTableRow(os.Args[1], os.Args[2], os.Args[3])
   fmt.Printf("%#v\n", row)
   umber := os.Getenv("UMBER")
   var rows []tableRow
   file, err := os.Open(umber)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   json.NewDecoder(file).Decode(&rows)
   rows = append([]tableRow{row}, rows...)
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
