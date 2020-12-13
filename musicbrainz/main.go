package main

import (
   "fmt"
   "log"
   "musicbrainz/release"
   "os"
   "path"
   "strings"
   "time"
)

func main() {
   if len(os.Args) != 2 {
      println(`Usage:
musicbrainz-release <URL>

URL:
https://musicbrainz.org/release/7cc21f46-16b4-4479-844c-e779572ca834
https://musicbrainz.org/release-group/67898886-90bd-3c37-a407-432e3680e872`)
      os.Exit(1)
   }
   url_s := os.Args[1]
   mbid_s := path.Base(url_s)
   dec_o := decode.NewDecode(mbid_s)
   rel_m := release.Map{}
   if strings.Contains(url_s, "release-group") {
      rel_a, e := dec_o.Group()
      if e != nil {
         log.Fatal(e)
      }
      rel_n := 0
      for idx_n, cur_m := range rel_a {
         rel_n = release.Reduce(rel_n, cur_m, idx_n, rel_a)
      }
      rel_m = rel_a[rel_n]
   } else {
      rel_m = dec_o.Release()
   }
   min_n := 179.5 * time.Second
   max_n := 15 * time.Minute
   album_m := map[string]string{
      "@date": rel_m["date"],
   }
   for _, media_m := range rel_m["media"] {
      for _, track_m := range media_m["tracks"] {
         len_n := track_m["length"] * Millisecond
         note_s := ""
         if len_n < min_n {
            note_s = "short"
         }
         if len_n > max_n {
            note_s = "long"
         }
         album_m[track_m["title"]] = note_s
      }
   }
   fmt.Println(album_m)
}
