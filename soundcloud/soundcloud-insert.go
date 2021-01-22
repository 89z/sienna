package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "path"
   "strconv"
)

func encode36(x int) string {
   n := int64(x)
   return strconv.FormatInt(n, 36)
}

func findSubmatch(pat string, sub []byte) []byte {
   a := regexp.MustCompile(pat).FindSubmatch(sub)
   if len(a) < 2 {
      return []byte{}
   }
   return a[1]
}

func main() {
   if len(os.Args) != 2 {
      println("soundcloud-insert <URL>")
      os.Exit(1)
   }
   url := os.Args[1]
   dest := path.Base(url) + ".html"
   if ! x.IsFile(dest) {
      _, e := x.HttpCopy(url, dest)
      if e != nil {
         log.Fatal(e)
      }
   }
   y, e := ioutil.ReadFile(dest)
   if e != nil {
      log.Fatal(e)
   }
   /*
   preg_match('!/tracks/([^"]*)"!', $dec_s, $track_a);
   $dec_s = html_entity_decode($get_s);
   $audio_s = $track_a[1];
   preg_match('!/artworks-([^.]*)-t500x500\.!', $dec_s, $img_a);
   $video_s = $img_a[1];
   preg_match('/ pubdate>(\d{4})-/', $dec_s, $year_a);
   $year_s = $year_a[1];
   $year_n = (int)($year_s);
   preg_match('/<title>([^|]*) by ([^|]*) \|/', $dec_s, $title_a);
   $title_s = $title_a[2] . ' - ' . $title_a[1];
   $date_n = time();
   $date_s = encode36($date_n);
   $rec_a = [$date_s, $year_n, 's/' . $audio_s . '/' . $video_s, $title_s];
   $json_s = json_encode($rec_a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
   echo $json_s, ",\n";
   */
}
