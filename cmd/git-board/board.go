package main

import (
   "fmt"
   "github.com/89z/rosso"
   "os"
   "strings"
   "time"
)

const minimum = 64

type board struct {
   actual, target string
   totAdd, totCha, totDel int
}

func newBoard() (board, error) {
   var cmd rosso.Cmd
   cmd.Run("git", "add", ".")
   arg := []string{"diff", "--cached", "--numstat"}
   _, e := os.Stat("config.toml")
   if e == nil {
      arg = append(arg, ":!docs")
   }
   stat, e := cmd.Out("git", arg...)
   if e != nil {
      return board{}, e
   }
   var b board
   for _, each := range strings.Split(stat, "\n") {
      b.totCha++
      var add, del int
      fmt.Sscan(each, &add, &del)
      b.totAdd += add
      b.totDel += del
   }
   commit, e := cmd.Out("git", "log", "-1", "--format=%cI")
   if e != nil {
      return board{}, e
   }
   b.actual = commit[:10]
   b.target = time.Now().AddDate(0, 0, -1).String()[:10]
   return b, nil
}

func main() {
   b, e := newBoard()
   if e != nil {
      panic(e)
   }
   for _, each := range []struct{
      name string
      actual, target interface{}
      result bool
   } {
      {"additions", b.totAdd, minimum, b.totAdd >= minimum},
      {"deletions", b.totDel, minimum, b.totDel >= minimum},
      {"changed files", b.totCha, minimum, b.totCha >= minimum},
      {"last commit", b.actual, b.target, b.actual <= b.target},
   } {
      message := fmt.Sprintf(
         "%-16v target: %-12v actual: %v", each.name, each.target, each.actual,
      )
      if each.result {
         rosso.LogPass("Pass", message)
      } else {
         rosso.LogFail("Fail", message)
      }
   }
}
