package main

import (
   "bufio"
   "net/url"
   "os"
   "path"
   "strings"
)

var mirror = url.URL{Scheme: "http", Host: "repo.msys2.org"}

// baseName(line, "=>")
func baseName(s, chars string) string {
   n := strings.IndexAny(s, chars)
   if n == -1 {
      return s
   }
   return s[:n]
}

type database struct {
   // %NAME% -> %FILENAME%, %DEPENDS%
   name map[string]description
   // %PROVIDES% -> %NAME%
   provides map[string]string
}

func newDatabase() database {
   return database{
      map[string]description{}, map[string]string{},
   }
}

func (db database) scan(file string) error {
   open, e := os.Open(file)
   if e != nil {
      return e
   }
   defer open.Close()
   scan := bufio.NewScanner(open)
   var filename, name string
   for scan.Scan() {
      switch scan.Text() {
      case "%FILENAME%":
         scan.Scan()
         filename = path.Join(repo, variant, scan.Text())
      case "%NAME%":
         scan.Scan()
         name = scan.Text()
      case "%PROVIDES%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" { break }
            db.provides[line] = name
         }
      case "%DEPENDS%":
         desc := description{filename: filename}
         for scan.Scan() {
            line := scan.Text()
            if line == "" { break }
            desc.depends = append(desc.depends, line)
         }
         db.name[name] = desc
      }
   }
   return nil
}

type description struct {
   filename string
   depends []string
}

type pack struct {
   repo string
   variant string
   file string
}

func (p pack) join(s string) string {
   return path.Join(s, p.repo, p.variant, p.file)
}
