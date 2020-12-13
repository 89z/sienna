package sql

import (
   "database/sql"
   "encoding/json"
   "log"
   "musicdb/assert"
   "os"
   _ "github.com/mithrandie/csvq-driver"
)

func Create() {
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
   res_o, e := sql_o.Exec("create table `artist_t.csv` (artist_n, artist_s)")
   if e != nil {
      log.Fatal(res_o, e)
   }
}
