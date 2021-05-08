package main

import (
   "fmt"
   "github.com/89z/rosso"
   "os"
   "os/exec"
   "strings"
   "time"
)

const minimum = 64

type board struct {
   actual, target string
   totAdd, totCha, totDel int
}

func newBoard() (board, error) {
   {
      cmd := exec.Command("git", "add", ".")
      fmt.Println("Run", cmd)
      cmd.Run()
   }
   arg := []string{"diff", "--cached", "--numstat"}
   _, err := os.Stat("config.toml")
   if err == nil {
      arg = append(arg, ":!docs")
   }
   var cmd rosso.Cmd
   stat, err := cmd.Out("git", arg...)
   if err != nil {
      return board{}, err
   }
   var b board
   /*
   7       5       cmd/fs-iterate/fs-iterate.go
   20      15      cmd/git-board/board.go
   9       11      cmd/youtube-insert/insert.go
   */
   for _, line := range strings.Split(stat, "\n") {
      b.totCha++
      var add, del int
      fmt.Sscan(line, &add, &del)
      b.totAdd += add
      b.totDel += del
   }
   commit, err := cmd.Out("git", "log", "-1", "--format=%cI")
   if err != nil {
      return board{}, err
   }
   b.actual = commit[:10]
   b.target = time.Now().AddDate(0, 0, -1).String()[:10]
   return b, nil
}

func main() {
   b, err := newBoard()
   if err != nil {
      panic(err)
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
