package main

import (
   "encoding/json"
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "regexp"
   "strconv"
   "time"
)

func findSubmatch(re string, input []byte) string {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return string(a[1])
}

func insert(url string) (string, error) {
   get, e := http.Get(url)
   if e != nil {
      return "", e
   }
   input, e := ioutil.ReadAll(get.Body)
   if e != nil {
      return "", e
   }
   audio := findSubmatch(`/tracks/([^"]*)"`, input)
   title := findSubmatch("<title>([^|]+) by ", input)
   video := findSubmatch("-([^-]+-[^-]+)-t500x500.jpg", input)
   year_s := findSubmatch(` pubdate>(\d{4})-`, input)
   year_n, e := strconv.Atoi(year_s)
   if e != nil {
      return "", e
   }
   date_n := time.Now().Unix()
   date_s := strconv.FormatInt(date_n, 36)
   rec, e := json.Marshal([]interface{}{
      date_s, year_n, "s/" + audio + "/" + video, title,
   })
   if e != nil {
      return "", e
   }
   return string(rec), nil
}

func main() {
   if len(os.Args) != 2 {
      println("soundcloud-insert <URL>")
      os.Exit(1)
   }
   url := os.Args[1]
   rec, e := insert(url)
   if e != nil {
      log.Fatal(e)
   }
   println(rec)
}
