package main

import (
   "database/sql"
   "fmt"
   "log"
   "os"
   _ "github.com/mattn/go-sqlite3"
)

func main() {
   if len(os.Args) < 3 {
      fmt.Println(`Synopsis:
   musicdb <target> <flags>

Examples:
   musicdb artist 'Kate Bush'
   musicdb artist 999 2019-12-31
   musicdb album 999 youtube.com/watch?v=HQmmM_qwG4k
   musicdb song 999 good`)
      os.Exit(1)
   }
   db_s := os.Getenv("MUSICDB")
   open_o, e := sql.Open("sqlite3", db_s)
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
   if os.Args[1] == "album" {
      album_s, url_s := os.Args[2], os.Args[3]
      e = AlbumUpdate(open_o, album_s, url_s)
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
   artist_s := os.Args[2]
   e = ArtistSelect(open_o, artist_s)
   if e != nil {
      log.Fatal(e)
   }
}
