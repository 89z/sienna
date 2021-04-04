package main

import (
   "encoding/json"
   "fmt"
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

func httpHead(addr string) bool {
   x.LogInfo("Head", addr)
   resp, e := http.Head(addr)
   return e == nil && resp.StatusCode == 200
}

func newTableRow(enc string) (tableRow, error) {
   dec, e := url.Parse(enc)
   if e != nil {
      return tableRow{}, e
   }
   id := dec.Query().Get("v")
   // year
   info, e := youtube.Info(id)
   if e != nil {
      return tableRow{}, e
   }
   if info.Description.SimpleText == "" {
      return tableRow{}, fmt.Errorf("Clapham Junction")
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
   val := make(url.Values)
   val.Set("a", now)
   val.Set("b", id)
   val.Set("p", "y")
   val.Set("y", year)
   switch {
   case httpHead("http://i.ytimg.com/vi/" + id + "/sddefault.jpg"):
   case httpHead("http://i.ytimg.com/vi/" + id + "/sd1.jpg"):
      val.Set("c", "sd1")
   default:
      val.Set("c", "hqdefault")
   }
   return tableRow{
      val.Encode(), title,
   }, nil
}

type tableRow struct { Q, S string }

func main() {
   if len(os.Args) != 2 {
      fmt.Println("youtube-insert <URL>")
      os.Exit(1)
   }
   // decode
   umber := os.Getenv("UMBER")
   file, e := os.Open(umber)
   if e != nil {
      log.Fatal(e)
   }
   var rows []tableRow
   json.NewDecoder(file).Decode(&rows)
   // append
   arg := os.Args[1]
   row, e := newTableRow(arg)
   if e != nil {
      log.Fatal(e)
   }
   rows = append([]tableRow{row}, rows...)
   // encode
   file, e = os.Create(umber)
   if e != nil {
      log.Fatal(e)
   }
   enc := json.NewEncoder(file)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(rows)
}
