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
   musicdb 'Kate Bush'
   musicdb check 999 2019-12-31
   musicdb note 999 good
   musicdb pop 999 0
   musicdb url 999 youtube.com/watch?v=HQmmM_qwG4k`)
      os.Exit(1)
   }
   db_s := os.Getenv("MUSICDB")
   open_o, e := sql.Open("sqlite3", db_s)
   if e != nil {
      log.Fatal(e)
   }
   key_s := os.Args[1]
   switch key_s {
   case "check":
      artist_s, check_s := os.Args[2], os.Args[3]
      e = CheckUpdate(open_o, artist_s, check_s)
   case "note":
      song_s, note_s := os.Args[2], os.Args[3]
      e = NoteUpdate(open_o, song_s, note_s)
   case "pop":
      println("pop")
   case "url":
      album_s, url_s := os.Args[2], os.Args[3]
      e = UrlUpdate(open_o, album_s, url_s)
   default:
      e = ArtistSelect(open_o, key_s)
   }
   if e != nil {
      log.Fatal(e)
   }
}
