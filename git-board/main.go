package main

import (
   "fmt"
   "github.com/89z/x"
   "strings"
   "time"
)

func main() {
   e := x.System("git", "add", ".")
   x.Check(e)
   stat, e := diff()
   x.Check(e)
   for stat.Scan() {
      totCha++
      text := stat.Text()
      if strings.HasPrefix(text, "-") {
         continue
      }
      fmt.Sscanf(text, "%v\t%v", &add, &del)
      totAdd += add
      totDel += del
   }
   commit, e := popen("git", "log", "--format=%cI")
   x.Check(e)
   commit.Scan()
   then := commit.Text()[:10]
   now := time.Now().String()[:10]
   for _, each := range []test{
      {"additions", totAdd, minimum, totAdd >= minimum},
      {"deletions", totDel, minimum, totDel >= minimum},
      {"changed files", totCha, minimum, totCha >= minimum},
      {"last commit date", then, now, now > then},
   } {
      message := fmt.Sprint(
         each.name, ": ", each.actual, ", target: ", each.target,
      )
      if each.result {
         fmt.Println(x.ColorGreen(message))
      } else {
         fmt.Println(x.ColorRed(message))
      }
   }
}
