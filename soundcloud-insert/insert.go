package main

import (
   "github.com/89z/x"
   "github.com/89z/x/soundcloud"
   "net/url"
   "os"
   "strconv"
   "time"
)

func main() {
   if len(os.Args) != 2 {
      println("soundcloud-insert <URL>")
      os.Exit(1)
   }
   player, e := soundcloud.Insert(os.Args[1])
   x.Check(e)
   value := make(url.Values)
   date := strconv.FormatInt(
      time.Now().Unix(), 36,
   )
   value.Set("a", date)
   value.Set("b", player.Id)
   value.Set("c", player.Artwork)
   value.Set("p", "s")
   value.Set("y", player.Pubdate)
   rec, e := x.JsonMarshal(map[string]string{
      "q": value.Encode(), "s": player.Title,
   })
   x.Check(e)
   os.Stdout.Write(append(rec, ',', '\n'))
}
