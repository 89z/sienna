package main

import (
   "bufio"
   "bytes"
   "log"
   "os"
   "os/exec"
   "strconv"
   "strings"
   "time"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func output(command ...string) (*bufio.Scanner, error) {
   name, arg := command[0], command[1:]
   b, e := exec.Command(name, arg...).Output()
   return bufio.NewScanner(bytes.NewReader(b)), e
}

func scanLines(data []byte, EOF bool) (int, []byte, error) {
   if EOF {
      return 0, nil, nil
   }
   if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
      return i + 2, data[:i], nil
   }
   return len(data), data[:len(data) - 1], nil
}

func touch(filename string, sec int64) error {
   t := time.Unix(sec, 0)
   return os.Chtimes(filename, t, t)
}

func main() {
   gitLs, e := output("git", "ls-files")
   check(e)
   file := map[string]bool{}
   for gitLs.Scan() {
      file[gitLs.Text()] = true
   }
   gitLog, e := output(
      "git", "log", "--name-only", "--relative", "--pretty=format:%ct", ".",
   )
   check(e)
   gitLog.Split(scanLines)
   for len(file) > 0 {
      gitLog.Scan()
      commit := strings.Split(gitLog.Text(), "\n")
      unix, e := strconv.ParseInt(commit[0], 10, 64)
      check(e)
      for _, name := range commit[1:] {
         if ! file[name] {
            continue
         }
         println(unix, "\t", name)
         e = touch(name, unix)
         check(e)
         delete(file, name)
      }
   }
}
