package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "strconv"
   "time"
)

func main() {
   gitLs, e := x.Popen("git", "ls-files")
   if e != nil {
      log.Fatal(e)
   }
   files := map[string]bool{}
   for gitLs.Scan() {
      files[gitLs.Text()] = true
   }
   gitLog, e := x.Popen(
      "git", "log", "-m",
      "--name-only", "--relative", "--pretty=format:%ct", ".",
   )
   if e != nil {
      log.Fatal(e)
   }
   for len(files) > 0 {
      gitLog.Scan()
      sec, e := strconv.ParseInt(gitLog.Text(), 10, 64)
      if e != nil {
         log.Fatal(e)
      }
      unix := time.Unix(sec, 0)
      for gitLog.Scan() {
         name := gitLog.Text()
         if name == "" {
            break
         }
         if ! files[name] {
            continue
         }
         println(sec, "\t", name)
         os.Chtimes(name, unix, unix)
         delete(files, name)
      }
   }
}
