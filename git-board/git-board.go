package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "log"
   "strings"
   "time"
)

const min = 64
var layout = time.RFC3339[:10]

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

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
   check(e)
   stat, e := diff()
   check(e)
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
   fmt.Println("minimum:", min)
   color(totCha >= min, "changed files", totCha)
   color(totAdd >= min, "additions", totAdd)
   color(totDel >= min, "deletions", totDel)
   fmt.Println()
   commit, e := x.Popen("git", "log", "--format=%cI")
   check(e)
   commit.Scan()
   parse, e := time.Parse(time.RFC3339, commit.Text())
   check(e)
   then := parse.Format(layout)
   now := time.Now().Format(layout)
   fmt.Println("last commit date:", then)
   color(then != now, "current date", now)
}
