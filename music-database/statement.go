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
      date_prev_s string
      song_n int
      song_s string
      note_s string
      artist_n int
      check_s string
   )
   for query_o.Next() {
      e = query_o.Scan(
         &album_s, &date_s, &song_n, &song_s, &note_s, &artist_n, &check_s,
      )
      if e != nil {
         return e
      }
      if date_s != date_prev_s {
         if date_prev_s != "" {
            fmt.Println()
         }
         th_s := date_s + " " + album_s
         hr_s := strings.Repeat("-", len(th_s))
         fmt.Print(th_s, "\n", hr_s, "\n")
         date_prev_s = date_s
      }
      // print space
      note_next_s := note_s
      if note_s == "" {
         note_next_s = fmt.Sprint(song_n)
      }
      fmt.Print(strings.Repeat(" ", 7 - len(note_next_s)))
      // print note
      if note_s == "" {
         fmt.Print("\x1b[30;43m", song_n, "\x1b[m")
      } else {
         fmt.Print(note_s)
      }
      // print song
      fmt.Println(" |", song_s)
   }
   if check_s == "" {
      check_s = fmt.Sprint("\x1b[30;43m", artist_n, "\x1b[m")
   }
   fmt.Print("\ncheck: ", check_s, "\n")
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
