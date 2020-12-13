package main

import (
   "database/sql"
   "log"
   _ "github.com/mithrandie/csvq-driver"
)

func main() {
   open_o, e := sql.Open("csvq", ".")
   if e != nil {
      log.Fatal(e)
   }
   _, e = open_o.Exec("SET @@ANSI_QUOTES TO TRUE")
   if e != nil {
      log.Fatal(e)
   }
   _, e = open_o.Exec(`CREATE TABLE "artist_t.csv" (artist_n, artist_s)`)
   if e != nil {
      log.Fatal(e)
   }
}
