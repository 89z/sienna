package main

import (
   "bufio"
   "github.com/89z/x"
   "os/exec"
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

type test struct {
   name string
   actual interface{}
   target interface{}
   result bool
}

var add, del, totAdd, totCha, totDel int
