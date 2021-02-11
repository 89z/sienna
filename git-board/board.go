package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "os"
   "strings"
   "time"
)

const minimum = 64
var add, del, totAdd, totCha, totDel int

func diff() (*bufio.Scanner, error) {
   _, e := os.Stat("config.toml")
   if e != nil {
      return popen("git", "diff", "--cached", "--numstat")
   }
   return popen("git", "diff", "--cached", "--numstat", ":!docs")
}

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := x.Command(name, arg...)
   pipe, e := cmd.StdoutPipe()
   if e != nil {
      return nil, fmt.Errorf("StdoutPipe %v", e)
   }
   return bufio.NewScanner(pipe), cmd.Start()
}

type test struct {
   name string
   actual interface{}
   target interface{}
   result bool
}

func main() {
   e := x.Command("git", "add", ".").Run()
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
   // actual
   actual := commit.Text()[:10]
   // target
   target := time.Now().AddDate(0, 0, -1).String()[:10]
   // print
   for _, each := range []test{
      {"additions", totAdd, minimum, totAdd >= minimum},
      {"deletions", totDel, minimum, totDel >= minimum},
      {"changed files", totCha, minimum, totCha >= minimum},
      {"last commit", actual, target, actual <= target},
   } {
      message := fmt.Sprintf(
         "%-16v target: %-12v actual: %v", each.name, each.target, each.actual,
      )
      if each.result {
         fmt.Println(x.ColorGreen(message))
      } else {
         fmt.Println(x.ColorRed(message))
      }
   }
}
