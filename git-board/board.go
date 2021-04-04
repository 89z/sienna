package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
   "strings"
   "time"
)

const minimum = 64

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := exec.Command(name, arg...)
   pipe, e := cmd.StdoutPipe()
   if e != nil {
      return nil, fmt.Errorf("StdoutPipe %v", e)
   }
   return bufio.NewScanner(pipe), cmd.Start()
}

func newBoard() (board, error) {
   exec.Command("git", "add", ".").Run()
   arg := []string{"diff", "--cached", "--numstat"}
   _, e := os.Stat("config.toml")
   if e == nil {
      arg = append(arg, ":!docs")
   }
   stat, e := popen("git", arg...)
   if e != nil {
      return board{}, e
   }
   var b board
   for stat.Scan() {
      b.totCha++
      text := stat.Text()
      if strings.HasPrefix(text, "-") { continue }
      var add, del int
      fmt.Sscan(text, &add, &del)
      b.totAdd += add
      b.totDel += del
   }
   commit, e := popen("git", "log", "--format=%cI")
   if e != nil {
      return board{}, e
   }
   commit.Scan()
   b.actual = commit.Text()[:10]
   b.target = time.Now().AddDate(0, 0, -1).String()[:10]
   return b, nil
}

type board struct {
   actual, target string
   totAdd, totCha, totDel int
}

func main() {
   b, e := newBoard()
   if e != nil {
      log.Fatal(e)
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
         x.LogPass("Pass", message)
      } else {
         x.LogFail("Fail", message)
      }
   }
}






