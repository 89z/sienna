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
         //filename = path.Join(repo, variant, scan.Text())
         filename = scan.Text()
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

type install struct {
   cache string
   destination string
}

var a = []string{
   "http://repo.msys2.org/msys/x86_64/msys.db.tar.gz",
   `%LocalAppData%\sienna\msys2\msys\x86_64\msys.db.tar.gz`,
   `%LocalAppData%\sienna\msys2`,
}

var b = []string{
   "http://repo.msys2.org/msys/x86_64/zstd-1.4.8-1-x86_64.pkg.tar.zst",
   `%LocalAppData%\sienna\msys2\msys\x86_64\zstd-1.4.8-1-x86_64.pkg.tar.zst`,
   `C:\sienna\msys2`,
}
