package main

import (
   "database/sql"
   "fmt"
   "strings"
   _ "github.com/mattn/go-sqlite3"
)

func ArtistSelect(open_o *sql.DB, artist_s string) error {
   query_s := `
   SELECT
      album_n, album_s, date_s, url_s,
      song_n, song_s, note_s,
      artist_n, check_s
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
      album_n int
      album_s string
      date_s string
      date_prev_s string
      url_s string
      song_n int
      song_s string
      note_s string
      artist_n int
      check_s string
   )
   for query_o.Next() {
      e = query_o.Scan(
         &album_n,
         &album_s,
         &date_s,
         &url_s,
         &song_n,
         &song_s,
         &note_s,
         &artist_n,
         &check_s,
      )
      if e != nil {
         return e
      }
      if date_s != date_prev_s {
         if date_prev_s != "" {
            fmt.Println()
         }
         // print album date, title
         fmt.Print(date_s, "\n", album_s, "\n")
         // print URL
         if url_s != "" {
            fmt.Println(url_s)
         } else {
            fmt.Print("\x1b[30;43m", album_n, "\x1b[m\n")
         }
         // print rule
         fmt.Println(strings.Repeat("-", 30))
         date_prev_s = date_s
      }
      note_s, color_b := SongNote(note_s, url_s, song_n)
      // print song space
      fmt.Printf("%*v", 7 - len(note_s), " ")
      // print song note
      if color_b {
         fmt.Print("\x1b[30;43m", note_s, "\x1b[m")
      } else {
         fmt.Print(note_s)
      }
      // print song title
      fmt.Println(" |", song_s)
   }
   // print artist check
   fmt.Print("\ncheck: ")
   if check_s != "" {
      fmt.Println(check_s)
   } else {
      fmt.Print("\x1b[30;43m", artist_n, "\x1b[m\n")
   }
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

func AlbumUpdate(open_o *sql.DB, album_s, url_s string) error {
   query_s := `
   UPDATE album_t SET url_s = ?
   WHERE album_n = ?
   `
   exec_o, e := open_o.Exec(query_s, url_s, album_s)
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
