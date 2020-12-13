package main

import (
   "database/sql"
   "encoding/json"
   "log"
   "musicdb/assert"
   "os"
   _ "github.com/mithrandie/csvq-driver"
)

func main() {
   os_o, e := os.Open("[C].json")
   if e != nil {
      log.Fatal(e)
   }
   json_m := assert.Map{}
   e = json.NewDecoder(os_o).Decode(&json_m)
   if e != nil {
      log.Fatal(e)
   }
   sql_o, e := sql.Open("csvq", "")
   if e != nil {
      log.Fatal(e)
   }
   for artist_s := range json_m {
      s := `insert into artist_t values (
         (select coalesce(max(artist_n), 0) + 1 from artist_t), ?
      )`
      res_o, e := sql_o.Exec(s, artist_s)
      if e != nil {
         log.Fatal(res_o, e)
      }
   }
}
