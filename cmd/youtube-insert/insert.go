package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/rosso"
   "github.com/89z/youtube"
   "net/http"
   "net/url"
   "os"
   "regexp"
   "strconv"
   "strings"
   "time"
)

func httpHead(addr string) bool {
   rosso.LogInfo("Head", addr)
   resp, err := http.Head(addr)
   return err == nil && resp.StatusCode == 200
}

func newTableRow(enc string) (tableRow, error) {
   dec, err := url.Parse(enc)
   if err != nil {
      return tableRow{}, err
   }
   id := dec.Query().Get("v")
   // year
   video, err := youtube.NewVideo(id)
   if err != nil {
      return tableRow{}, err
   }
   if video.Description() == "" {
      return tableRow{}, fmt.Errorf("Clapham Junction")
   }
   year := video.Microformat.PlayerMicroformatRenderer.PublishDate[:4]
   for _, pattern := range []string{
      /* the order doesnt matter here, as we will find the lowest date of all
      matches */
      ` (\d{4})`, `(\d{4}) `, `Released on: (\d{4})`, `℗ (\d{4})`,
   } {
      re := regexp.MustCompile(pattern)
      find := re.FindStringSubmatch(video.Description())
      if len(find) < 2 { continue }
      if find[1] >= year { continue }
      year = find[1]
   }
   // song, artist
   title := video.Title()
   line := regexp.MustCompile(".* · .*").FindString(video.Description())
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
   file, err := os.Open(umber)
   if err != nil {
      panic(err)
   }
   var rows []tableRow
   json.NewDecoder(file).Decode(&rows)
   // append
   arg := os.Args[1]
   row, err := newTableRow(arg)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", row)
   rows = append([]tableRow{row}, rows...)
   // encode
   file, err = os.Create(umber)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(rows)
}
