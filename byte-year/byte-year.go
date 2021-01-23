package main

import (
   "fmt"
   "github.com/89z/x"
   "log"
   "os"
)

const year_sec = 365 * 24 * 60 * 60

var (
   err error
   file *bufio.Scanner
)

type entry struct {
   name string
   size int
}

func main() {
   if len(os.Args) == 1 {
      file, err = x.Popen("git", "ls-files")
   } else {
      arg := os.Args[1]
      file, err = x.Popen("git", "ls-files", ":!" + arg)
   }
   files := []entry{}
   for file.Scan() {
      name := file.Text()
      then, e := x.ModTime(name)
      if e != nil {
         log.Fatal(e)
      }
      year := x.SinceHours(then) / 24 / 365
      files = append(files, entry{
         name,
         'size' => filesize($name_s) * $year_n,
      })
   }
   $f = fn ($m, $m2) => $m2['size'] <=> $m['size'];
   usort($file_a, $f);
   foreach ($file_a as $file_m) {
      echo $file_m['size'], ' ', $file_m['name'], "\n";
   }
}
