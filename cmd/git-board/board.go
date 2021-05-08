package main

import (
   "fmt"
   "os"
   "os/exec"
   "strings"
   "time"
)

const minimum = 64

const (
   reset = "\x1b[m"
   green = "\x1b[30;102m"
   invert = "\x1b[7m"
   red = "\x1b[30;101m"
)

type board struct {
   actual, target string
   totAdd, totCha, totDel int
}

func newBoard() (board, error) {
   cmd := exec.Command("git", "add", ".")
   fmt.Println(invert, "Run", reset, cmd)
   cmd.Run()
   arg := []string{"diff", "--cached", "--numstat"}
   _, err := os.Stat("config.toml")
   if err == nil {
      arg = append(arg, ":!docs")
   }
   cmd = exec.Command("git", arg...)
   pipe, err := cmd.StdoutPipe()
   if err != nil {
      return board{}, err
   }
   cmd.Start()
   defer cmd.Wait()
   var b board
   for {
      var stat struct {
         add, del int
         path string
      }
      _, err := fmt.Fscanln(pipe, &stat.add, &stat.del, &stat.path)
      if err != nil { break }
      b.totCha += 1
      b.totAdd += stat.add
      b.totAdd += stat.del
   }
   out := new(strings.Builder)
   cmd = exec.Command("git", "log", "-1", "--format=%cI")
   cmd.Stdout = out
   cmd.Run()
   b.actual = out.String()[:10]
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
         fmt.Println(green, "Pass", reset, message)
      } else {
         fmt.Println(red, "Fail", reset, message)
      }
   }
}
