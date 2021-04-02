package main

import (
   "github.com/89z/x"
   "github.com/89z/x/youtube"
   "log"
   "net/http"
   "net/url"
   "os"
   "regexp"
   "strconv"
   "strings"
   "time"
)

func getImage(id string) string {
   switch {
   case httpHead("http://i.ytimg.com/vi/" + id + "/sddefault.jpg"):
      return ""
   case httpHead("http://i.ytimg.com/vi/" + id + "/sd1.jpg"):
      return "sd1"
   default:
      return "hqdefault"
   }
}

func httpHead(s string) bool {
   x.LogInfo("Head", s)
   resp, e := http.Head(s)
   return e == nil && resp.StatusCode == 200
}

func main() {
   if len(os.Args) != 2 {
      println("youtube-insert <URL>")
      os.Exit(1)
   }
   arg := os.Args[1]
   addr, e := url.Parse(arg)
   if e != nil {
      log.Fatal(e)
   }
   id := addr.Query().Get("v")
   // year
   info, e := youtube.Info(id)
   if e != nil {
      log.Fatal(e)
   }
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
      re := regexp.MustCompile(pattern)
      find := re.FindStringSubmatch(info.Description.SimpleText)
      if len(find) < 2 { continue }
      if find[1] >= year { continue }
      year = find[1]
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
   now := strconv.FormatInt(time.Now().Unix(), 36)
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
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(append(data, ',', '\n'))
}
