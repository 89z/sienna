package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "os/exec"
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
