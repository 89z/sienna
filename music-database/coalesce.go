package main

import (
   "fmt"
   "strings"
)

func SongNote(note_s, url_s string, song_n int) (string, bool) {
   if note_s != "" {
      return note_s, false
   }
   if strings.HasPrefix(url_s, "youtu.be/") {
      return note_s, false
   }
   return fmt.Sprint(song_n), true
}
