package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "os"
   "sort"
   "time"
)

func lsFiles() (*bufio.Scanner, error) {
   if len(os.Args) == 1 {
      return x.Popen("git", "ls-files")
   }
   arg := os.Args[1]
   return x.Popen("git", "ls-files", ":!" + arg)
}

func main() {
   file, e := lsFiles()
   x.Check(e)
   files := []entry{}
   for file.Scan() {
      name := file.Text()
      then, e := x.ModTime(name)
      x.Check(e)
      size, e := x.FileSize(name)
      x.Check(e)
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
