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
   var add, del, totAdd, totCha, totDel int
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
   fmt.Println("minimum:", minimum)
   color(totCha >= minimum, "changed files", totCha)
   color(totAdd >= minimum, "additions", totAdd)
   color(totDel >= minimum, "deletions", totDel)
   fmt.Println()
   commit, e := popen("git", "log", "--format=%cI")
   x.Check(e)
   commit.Scan()
   then := commit.Text()[:10]
   now := time.Now().String()[:10]
   fmt.Println("last commit date:", then)
   color(then != now, "current date", now)
}
