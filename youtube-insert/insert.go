package main

import (
   "bytes"
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

func getImage(id string) string {
   url := "https://i.ytimg.com/vi/"
   if httpHead(url + id + "/sddefault.jpg") {
      return ""
   }
   if httpHead(url + id + "/sd1.jpg") {
      return "sd1"
   }
   return "hqdefault"
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
   if info.Description.SimpleText == "" {
      println("Clapham Junction")
      os.Exit(1)
   }
   year := info.PublishDate[:4]
   for _, pattern := range patterns {
      match := findSubmatch(pattern, info.Description.SimpleText)
      if match == "" {
         continue
      }
      if match >= year {
         continue
      }
      year = match
   }
   // song, artist
   title := info.Title.SimpleText
   line := regexp.MustCompile(".* · .*").FindString(info.Description.SimpleText)
   if line != "" {
      titles := strings.Split(line, " · ")
      artists := titles[1:]
      title = strings.Join(artists, ", ") + " - " + titles[0]
   }
   // time
   now := strconv.FormatInt(
      time.Now().Unix(), 36,
   )
   // print
   value := make(url.Values)
   value.Set("a", now)
   value.Set("b", id)
   value.Set("p", "y")
   value.Set("y", year)
   image := getImage(id)
   if image != "" {
      value.Set("c", image)
   }
   data, e := marshal(map[string]string{
      "q": value.Encode(), "s": title,
   })
   x.Check(e)
   os.Stdout.Write(append(data, ',', '\n'))
}
