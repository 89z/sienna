package main

import (
   "database/sql"
   "fmt"
   "log"
   "os"
   _ "github.com/mattn/go-sqlite3"
)

func ArtistSelect(open_o *sql.DB, artist_s string) error {
   query_s := `
   SELECT
      album_s, date_s,
      song_n, song_s, note_s,
      artist_n, COALESCE(check_s, '')
   FROM album_t
   NATURAL JOIN song_album_t
   NATURAL JOIN song_t
   NATURAL JOIN song_artist_t
   NATURAL JOIN artist_t
   WHERE artist_s = ?
   ORDER BY date_s
   `
   query_o, e := open_o.Query(query_s, artist_s)
   if e != nil {
      return e
   }
   var (
      album_s string
      date_s string
      song_n int
      song_s string
      note_s string
      artist_n int
      check_s string
   )
   for query_o.Next() {
      e := query_o.Scan(
         &album_s, &date_s, &song_n, &song_s, &note_s, &artist_n, &check_s,
      )
      if e != nil {
         return e
      }
      if note_s == "" {
         note_s = fmt.Sprint("\x1b[36m", song_n, "\x1b[m")
      }
      fmt.Printf("%.4v | %v | %v | %v\n", date_s, album_s, song_s, note_s)
   }
   if check_s == "" {
      check_s = fmt.Sprint("\x1b[36m", artist_n, "\x1b[m")
   }
   fmt.Println("check:", check_s)
   return nil
}

func ArtistUpdate(open_o *sql.DB, artist_s, check_s string) error {
   query_s := `
   UPDATE artist_t SET check_s = ?
   WHERE artist_n = ?
   `
   exec_o, e := open_o.Exec(query_s, check_s, artist_s)
   if e != nil {
      return fmt.Errorf("%v %v", exec_o, e)
   }
   return nil
}

func SongUpdate(open_o *sql.DB, song_s, note_s string) error {
   query_s := `
   UPDATE song_t SET note_s = ?
   WHERE song_n = ?
   `
   exec_o, e := open_o.Exec(query_s, note_s, song_s)
   if e != nil {
      return fmt.Errorf("%v %v", exec_o, e)
   }
   return nil
}

func main() {
   if len(os.Args) < 3 {
      fmt.Println(`Synopsis:
   musicdb <target> <flags>

Examples:
   musicdb artist 'Kate Bush'
   musicdb artist 30 2020-12-14
   musicdb song 10656 good`)
      os.Exit(1)
   }
   open_o, e := sql.Open("sqlite3", `D:\Music\music.db`)
   if e != nil {
      log.Fatal(e)
   }
   if os.Args[1] == "song" {
      song_s, note_s := os.Args[2], os.Args[3]
      e = SongUpdate(open_o, song_s, note_s)
      if e != nil {
         log.Fatal(e)
      }
      return
   }
   if len(os.Args) == 4 {
      artist_s, check_s := os.Args[2], os.Args[3]
      e = ArtistUpdate(open_o, artist_s, check_s)
      if e != nil {
         log.Fatal(e)
      }
      return
   }
   artist_s := os.Args[1]
   e = ArtistSelect(open_o, artist_s)
   if e != nil {
      log.Fatal(e)
   }
}
