package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "strings"
   "time"
)

const minimum = 64

func color(test bool, key string, value interface{}) {
   message := fmt.Sprint(key, ": ", value)
   if test {
      fmt.Println(x.ColorGreen(message))
   } else {
      fmt.Println(x.ColorRed(message))
   }
}

func diff() (*bufio.Scanner, error) {
   if x.IsFile("config.toml") {
      return x.Popen("git", "diff", "--cached", "--numstat", ":!docs")
   }
   return x.Popen("git", "diff", "--cached", "--numstat")
}

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
   commit, e := x.Popen("git", "log", "--format=%cI")
   x.Check(e)
   commit.Scan()
   then := commit.Text()[:10]
   now := time.Now().String()[:10]
   fmt.Println("last commit date:", then)
   color(then != now, "current date", now)
}
