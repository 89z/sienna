package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "os"
   "os/exec"
   "strings"
   "time"
)

const minimum = 64


func diff() (*bufio.Scanner, error) {
   if x.IsFile("config.toml") {
      return popen("git", "diff", "--cached", "--numstat", ":!docs")
   }
   return popen("git", "diff", "--cached", "--numstat")
}

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := exec.Command(name, arg...)
   pipe, err := cmd.StdoutPipe()
   if err != nil {
      return nil, err
   }
   return bufio.NewScanner(pipe), cmd.Start()
}

func board() error {
   e := x.System("git", "add", ".")
   if e != nil {
      return e
   }
   stat, e := diff()
   if e != nil {
      return e
   }
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
   commit, e := popen("git", "log", "--format=%cI")
   if e != nil {
      return e
   }
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
   return nil
}

type test struct {
   name string
   actual interface{}
   target interface{}
   result bool
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("git-train <board | depart>")
      os.Exit(1)
   }
   if os.Args[1] == "board" {
      e := board()
      x.Check(e)
      return
   }
   e := x.System("git", "commit", "--verbose")
   x.Check(e)
   if x.IsFile("config.toml") {
      fmt.Println(x.ColorGreen("remove docs"))
      os.RemoveAll("docs")
      fmt.Println(x.ColorGreen("hugo"))
      e = x.System("hugo")
      x.Check(e)
      fmt.Println(x.ColorGreen("git add"))
      e = x.System("git", "add", ".")
      x.Check(e)
      fmt.Println(x.ColorGreen("git commit"))
      e = x.System("git", "commit", "--amend")
      x.Check(e)
   }
}
