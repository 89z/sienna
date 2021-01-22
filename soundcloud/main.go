package main

import (
   "encoding/json"
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "os"
   "path"
   "regexp"
   "strconv"
   "time"
)

func main() {
   if len(os.Args) != 2 {
      println("soundcloud-insert <URL>")
      os.Exit(1)
   }
   url := os.Args[1]
   dest := path.Base(url) + ".html"
   if ! x.IsFile(dest) {
      _, e := x.HttpCopy(url, dest)
      check(e)
   }
   input, e := ioutil.ReadFile(dest)
   check(e)
   audio := findSubmatch(`/tracks/([^"]*)"`, input)
   video := findSubmatch("-([^-]+-[^-]+)-t500x500.jpg", input)
   title := findSubmatch("<title>([^|]+) by ", input)
   year_s := findSubmatch(` pubdate>(\d{4})-`, input)
   year_n, e := strconv.Atoi(year_s)
   check(e)
   date_n := time.Now().Unix()
   date_s := strconv.FormatInt(date_n, 36)
   rec := x.Slice{date_s, year_n, "s/" + audio + "/" + video, title}
   e = json.NewEncoder(os.Stdout).Encode(rec)
   check(e)
}
