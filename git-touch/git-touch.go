package main

import (
   "bufio"
   "log"
   "os/exec"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func popen(command ...string) (*bufio.Scanner, error) {
   name, arg := command[0], command[1:]
   cmd := exec.Command(name, arg...)
   pipe, e := cmd.StdoutPipe()
   if e != nil {
      return nil, e
   }
   return bufio.NewScanner(pipe), cmd.Start()
}
