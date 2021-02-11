package main

import (
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

func getImage(id string) string {
   if httpHead("http://i.ytimg.com/vi/" + id + "/sddefault.jpg") {
      return ""
   }
   if httpHead("http://i.ytimg.com/vi/" + id + "/sd1.jpg") {
      return "sd1"
   }
   return "hqdefault"
}

func httpHead(s string) bool {
   println(x.ColorCyan("Head"), s)
   resp, e := http.Head(s)
   return e == nil && resp.StatusCode == 200
}

func findStringSubmatch(re, input string) string {
   a := regexp.MustCompile(re).FindStringSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return a[1]
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
   for _, pattern := range []string{
      /* the order doesnt matter here, as we will find the lowest date of all
      matches */
      ` (\d{4})`, `(\d{4}) `, `Released on: (\d{4})`, `℗ (\d{4})`,
   } {
      match := findStringSubmatch(pattern, info.Description.SimpleText)
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
   data, e := x.JsonMarshal(map[string]string{
      "q": value.Encode(), "s": title,
   })
   x.Check(e)
   os.Stdout.Write(append(data, ',', '\n'))
}
