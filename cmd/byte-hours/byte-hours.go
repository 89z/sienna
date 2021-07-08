package main

import (
   "bufio"
   "fmt"
   "os"
   "os/exec"
   "sort"
   "time"
)

func fileSize(name string) (int64, error) {
   info, err := os.Stat(name)
   if err != nil {
      return 0, err
   }
   return info.Size(), nil
}

func modTime(name string) (time.Time, error) {
   i, err := os.Stat(name)
   if err != nil {
      return time.Time{}, err
   }
   return i.ModTime(), nil
}

func main() {
   cmd := exec.Command("git", "ls-files")
   pipe, err := cmd.StdoutPipe()
   if err != nil {
      panic(err)
   }
   cmd.Start()
   defer cmd.Wait()
   var (
      file = bufio.NewScanner(pipe)
      files []entry
   )
   for file.Scan() {
      name := file.Text()
      then, err := modTime(name)
      if err != nil {
         panic(err)
      }
      size, err := fileSize(name)
      if err != nil {
         panic(err)
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
