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

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func lsFiles() (*bufio.Scanner, error) {
   if len(os.Args) == 1 {
      return x.Popen("git", "ls-files")
   }
   arg := os.Args[1]
   return x.Popen("git", "ls-files", ":!" + arg)
}

func main() {
   file, e := lsFiles()
   check(e)
   files := []entry{}
   for file.Scan() {
      name := file.Text()
      then, e := x.ModTime(name)
      check(e)
      size, e := x.FileSize(name)
      check(e)
      year := time.Since(then).Hours() / 24 / 365
      files = append(files, entry{
         name, float64(size) * year,
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
   size float64
}
