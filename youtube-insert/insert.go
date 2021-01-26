package main

import (
   "encoding/json"
   "github.com/89z/x"
   "github.com/89z/x/youtube"
   "net/http"
   "net/url"
   "os"
   "regexp"
   "strconv"
   "strings"
   "time"
)

/* the order doesnt matter here, as we will find the lowest date of all
matches */
var patterns = []string{
   ` (\d{4})`, `(\d{4}) `, `Released on: (\d{4})`, `℗ (\d{4})`,
}

func findSubmatch(re, input string) string {
   a := regexp.MustCompile(re).FindStringSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}

func getImage(id string) string {
   url := "https://i.ytimg.com/vi/"
   if httpHead(url + id + "/sddefault.jpg") {
      return ""
   }
   if httpHead(url + id + "/sd1.jpg") {
      return "/sd1"
   }
   return "/hqdefault"
}

func httpHead(url string) bool {
   println(url)
   resp, e := http.Head(url)
   return e == nil && resp.StatusCode == 200
}

func main() {
   if len(os.Args) != 2 {
      println("youtube-insert <URL>")
      os.Exit(1)
   }
   u, e := url.Parse(os.Args[1])
   x.Check(e)
   id := u.Query().Get("v")
   // year
   info, e := youtube.Info(id)
   x.Check(e)
   if info["description"] == nil {
      println("Clapham Junction")
      os.Exit(1)
   }
   desc := info.M("description").S("simpleText")
   date := info.S("publishDate")[:4]
   for _, pattern := range patterns {
      match := findSubmatch(pattern, desc)
      if match == "" {
         continue
      }
      if match >= date {
         continue
      }
      date = match
   }
   year, e := strconv.Atoi(date)
   x.Check(e)
   // song, artist
   title := info.M("title").S("simpleText")
   line := regexp.MustCompile(".* · .*").FindString(desc)
   if line != "" {
      titles := strings.Split(line, " · ")
      artists := titles[1:]
      title = strings.Join(artists, ", ") + " - " + titles[0]
   }
   // time
   now := strconv.FormatInt(
      time.Now().Unix(), 36,
   )
   // image
   image := getImage(id)
   // print
   rec, e := json.Marshal(
      []interface{}{now, year, "y/" + id + image, title},
   )
   x.Check(e)
   os.Stdout.Write(append(rec, ',', '\n'))
}
