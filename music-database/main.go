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
   if len(os.Args) != 2 {
      println("musicdb <JSON>")
      os.Exit(1)
   }
   path_s := os.Args[1]
   os_o, e := os.Open(path_s)
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
      println(artist_s)
      s := `insert into artist_t values (
         (select coalesce(max(artist_n), 0) + 1 from artist_t), ?
      )`
      o, e := sql_o.Exec(s, artist_s)
      if e != nil {
         log.Fatal(o, e)
      }
      artist_m := json_m.M(artist_s)
      for album_s := range artist_m {
         s := `insert into album_t (album_n, album_s) values (
            (select coalesce(max(album_n), 0) + 1 from album_t), ?
         )`
         o, e := sql_o.Exec(s, album_s)
         if e != nil {
            log.Fatal(o, e)
         }
         album_m := artist_m.M(album_s)
         for song_s := range album_m {
            note_s := album_m.S(song_s)
            if song_s == "@date" {
               s := `update album_t set date_d = ?
               where album_n = (select max(album_n) from album_t)`
               o, e := sql_o.Exec(s, note_s)
               if e != nil {
                  log.Fatal(o, e)
               }
               continue
            }
            if song_s == "@url" {
               s := `update album_t set url_s = ?
               where album_n = (select max(album_n) from album_t)`
               o, e := sql_o.Exec(s, note_s)
               if e != nil {
                  log.Fatal(o, e)
               }
               continue
            }
            s := `insert into song_t values (
               (select coalesce(max(song_n), 0) + 1 from song_t), ?, ?
            )`
            o, e := sql_o.Exec(s, song_s, note_s)
            if e != nil {
               log.Fatal(o, e)
            }
            s = `insert into song_album_t values (
               (select max(song_n) from song_t),
               (select max(album_n) from album_t)
            )`
            o, e = sql_o.Exec(s)
            if e != nil {
               log.Fatal(o, e)
            }
            s = `insert into song_artist_t values (
               (select max(song_n) from song_t),
               (select max(artist_n) from artist_t)
            )`
            o, e = sql_o.Exec(s)
            if e != nil {
               log.Fatal(o, e)
            }
         }
      }
   }
}
