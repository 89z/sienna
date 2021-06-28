package main

import (
   "encoding/json"
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
   "math/big"
   "net/http"
   "net/url"
   "os"
   "path"
   "regexp"
   "sort"
   "strconv"
   "strings"
   "time"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func mod(a, b int64) int64 {
   c, d := big.NewInt(a), big.NewInt(b)
   return new(big.Int).Mod(c, d).Int64()
}

func score(i youtube.Image) int64 {
   return mod(480 - i.Height, 720) + i.Frame + i.Format
}

type tableRow struct {
   Q string
   S string
}

func main() {
   var dry bool
   flag.BoolVar(&dry, "d", false, "dry run")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("youtube-insert [flag] <URL>")
      flag.PrintDefaults()
      return
   }
   addr, err := url.Parse(flag.Arg(0))
   if err != nil {
      panic(err)
   }
   row, err := newTableRow(addr.Query().Get("v"))
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(row)
   if dry {
      return
   }
   // write
   umber := os.Getenv("UMBER")
   // decode
   var rows []*tableRow
   file, err := os.Open(umber)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   json.NewDecoder(file).Decode(&rows)
   // append
   rows = append([]*tableRow{row}, rows...)
   // encode
   if file, err := os.Create(umber); err != nil {
      panic(err)
   } else {
      defer file.Close()
      enc := json.NewEncoder(file)
      enc.SetEscapeHTML(false)
      enc.SetIndent("", " ")
      enc.Encode(rows)
   }
}


func newTableRow(id string) (*tableRow, error) {
   val := make(url.Values)
   val.Set("p", "y")
   val.Set("b", id)
   // image
   sort.SliceStable(youtube.Images, func(d, e int) bool {
      return score(youtube.Images[d]) < score(youtube.Images[e])
   })
   for n, img := range youtube.Images {
      addr := img.Address(id)
      fmt.Println(invert, "Head", reset, addr)
      res, err := http.Head(addr)
      if err == nil && res.StatusCode == http.StatusOK {
         if n > 0 {
            val.Set("c", path.Base(addr))
         }
         break
      }
   }
   // year
   video, err := youtube.NewMWeb(id)
   if err != nil {
      return nil, err
   }
   if video.ShortDescription == "" {
      return nil, fmt.Errorf("clapham Junction")
   }
   year := video.PublishDate[:4]
   for _, pat := range []string{
      /* the order doesnt matter here, as we will find the lowest date of all
      matches */
      ` (\d{4})`, `(\d{4}) `, `Released on: (\d{4})`, `℗ (\d{4})`,
   } {
      re := regexp.MustCompile(pat).FindStringSubmatch(video.ShortDescription)
      if re == nil {
         continue
      }
      if re[1] < year {
         year = re[1]
      }
   }
   val.Set("y", year)
   // song, artist
   re := regexp.MustCompile(".* · .*").FindString(video.ShortDescription)
   if re != "" {
      titles := strings.Split(re, " · ")
      artists := titles[1:]
      video.Title = strings.Join(artists, ", ") + " - " + titles[0]
   }
   // time
   now := strconv.FormatInt(time.Now().Unix(), 36)
   val.Set("a", now)
   // return
   return &tableRow{
      val.Encode(), video.Title,
   }, nil
}
