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
      album_s, date_s, url_s,
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
         len_n := len(album_s)
         if url_s != "" {
            len_n = len(url_s)
            fmt.Println(url_s)
         }
         // print rule
         hr_s := strings.Repeat("-", len_n)
         fmt.Println(hr_s)
         date_prev_s = date_s
      }
      // print song note, title
      note_s = SongNote(note_s, url_s, song_n)
      fmt.Println(note_s, "|", song_s)
   }
   if check_s == "" {
      check_s = fmt.Sprint("\x1b[30;43m", artist_n, "\x1b[m")
   }
   fmt.Print("\ncheck: ", check_s, "\n")
   return nil
}

func SongNote(note_s, url_s string, song_n int) string {
   if note_s != "" {
      return fmt.Sprintf("%7v", note_s)
   }
   if strings.HasPrefix(url_s, "youtu.be/") {
      return fmt.Sprintf("%7v", note_s)
   }
   return strings.Repeat(" ", 7 - len(fmt.Sprint(song_n))) +
   fmt.Sprint("\x1b[30;43m", song_n, "\x1b[m")
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
