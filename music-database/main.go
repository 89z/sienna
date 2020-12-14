package main

import (
   "database/sql"
   "log"
   "os"
   _ "github.com/mattn/go-sqlite3"
)

func Artist(open_o *sql.DB, artist_s string) error {
   query_s := `select song_s, note_s from song_t where song_n in (
      select song_n from song_artist_t where artist_n = (
         select artist_n from artist_t where artist_s = ?
      )
   )`
   query_o, e := open_o.Query(query_s, artist_s)
   if e != nil {
      return e
   }
   var song_s, note_s string
   for query_o.Next() {
      e := query_o.Scan(&song_s, &note_s)
      if e != nil {
         return e
      }
      println(song_s, "|", note_s)
   }
   return nil
}

func main() {
   if len(os.Args) != 2 {
      println("musicdb <artist>")
      os.Exit(1)
   }
   artist_s := os.Args[1]
   open_o, e := sql.Open("sqlite3", `D:\Music\music.db`)
   if e != nil {
      log.Fatal(e)
   }
   e = Artist(open_o, artist_s)
   if e != nil {
      log.Fatal(e)
   }
}
