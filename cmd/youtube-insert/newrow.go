package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "math"
   "net/http"
   "net/url"
   "path"
   "regexp"
   "sort"
   "strconv"
   "strings"
   "time"
)

type iFunc func(a, b youtube.Image) bool

var iFuncs = []iFunc{
   func(a, b youtube.Image) bool {
      return math.Copysign(1, a.Height-720) < math.Copysign(1, b.Height-720)
   },
   func(a, b youtube.Image) bool {
      return math.Abs(a.Height-720) < math.Abs(b.Height-720)
   },
   func(a, b youtube.Image) bool {
      return a.Frame < b.Frame
   },
   func(a, b youtube.Image) bool {
      return a.Format < b.Format
   },
}

func newTableRow(id string) (*tableRow, error) {
   val := make(url.Values)
   val.Set("p", "y")
   val.Set("b", id)
   // image
   sort.SliceStable(youtube.Images, func(a, b int) bool {
      ia, ib := youtube.Images[a], youtube.Images[b]
      for _, fn := range iFuncs {
         if fn(ia, ib) {
            return true
         }
         if fn(ib, ia) {
            break
         }
      }
      return false
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
