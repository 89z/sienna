package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "log"
   "os"
   "sort"
   "time"
)

func fileSize(name string) (int64, error) {
   info, e := os.Stat(name)
   if e != nil {
      return 0, e
   }
   return info.Size(), nil
}

func lsFiles() (*bufio.Scanner, error) {
   if len(os.Args) == 1 {
      return x.Popen("git", "ls-files")
   }
   arg := os.Args[1]
   return x.Popen("git", "ls-files", ":!" + arg)
}

func modTime(name string) (time.Time, error) {
   stat, e := os.Stat(name)
   if e != nil {
      return time.Time{}, e
   }
   return stat.ModTime(), nil
}
func main() {
   file, e := lsFiles()
   if e != nil {
      log.Fatal(e)
   }
   var files []entry
   for file.Scan() {
      name := file.Text()
      then, e := modTime(name)
      if e != nil {
         log.Fatal(e)
      }
      size, e := fileSize(name)
      if e != nil {
         log.Fatal(e)
      }
      hour := time.Since(then).Hours()
      files = append(files, entry{
         name, size * int64(hour),
      })
   }
   sort.Slice(files, func (i, j int) bool {
      return files[i].size < files[j].size
   })
   for _, each := range files {
      fmt.Println(each.size, each.name)
   }
}

type entry struct {
   name string
   size int64
}
