package main

import (
   "bufio"
   "os"
   "path"
   "strings"
)

const mirror = "http://repo.msys2.org"

func baseName(s, char string) string {
   n := strings.IndexAny(s, char)
   if n == -1 {
      return s
   }
   return s[:n]
}

type description struct {
   name string
   filename string
   provides []string
   depends []string
}

func newDescription(file, repo, variant string) (description, error) {
   open, e := os.Open(file)
   if e != nil {
      return description{}, e
   }
   scan := bufio.NewScanner(open)
   var desc description
   for scan.Scan() {
      switch scan.Text() {
      case "%FILENAME%":
         scan.Scan()
         desc.filename = path.Join(mirror, repo, variant, scan.Text())
      case "%NAME%":
         scan.Scan()
         desc.name = scan.Text()
      case "%DEPENDS%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" {
               break
            }
            desc.depends = append(
               desc.depends, baseName(line, "=>"),
            )
         }
      case "%PROVIDES%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" {
               break
            }
            desc.provides = append(desc.provides, line)
         }
      }
   }
   return desc, nil
}
