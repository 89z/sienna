package main

import (
   "database/sql"
   "log"
   _ "github.com/mattn/go-sqlite3"
)

var (
   artist_n int
   artist_s string
)

func main() {
   open_o, e := sql.Open("sqlite3", `D:\Music\music.db`)
   if e != nil {
      log.Fatal(e)
   }
   query_o, e := open_o.Query("select * from artist_t")
   if e != nil {
      log.Fatal(e)
   }
   for query_o.Next() {
      e := query_o.Scan(&artist_n, &artist_s)
      if e != nil {
         log.Fatal(e)
      }
      println(artist_n, artist_s)
   }
}
