package main

import (
   "bufio"
   "net/url"
   "os"
   "path/filepath"
   "strings"
)

var mirror = url.URL{Scheme: "http", Host: "repo.msys2.org"}

func baseName(s, chars string) string {
   n := strings.IndexAny(s, chars)
   if n == -1 {
      return s
   }
   return s[:n]
}

type database struct {
   // %NAME% -> %FILENAME%, %ARCH%, %DEPENDS%
   name map[string]description
   // %PROVIDES% -> %NAME%
   provides map[string]string
}

func newDatabase() database {
   return database{
      map[string]description{}, map[string]string{},
   }
}

func (db database) query(target string) {
   done := map[string]bool{target: true}
   for todo := []string{target}; len(todo) > 0; todo = todo[1:] {
      do := todo[0]
      for _, dep := range db.name[do].depends {
         if ! done[dep] {
            todo = append(todo, dep)
            done[dep] = true
         }
      }
      println(do)
   }
}

func (db database) scan(file []byte) error {
   var (
      desc description
      name string
      scan = bufio.NewScanner(bytes.NewReader(file))
   )
   for scan.Scan() {
      switch scan.Text() {
      case "%FILENAME%":
         scan.Scan()
         desc.filename = scan.Text()
      case "%NAME%":
         scan.Scan()
         name = scan.Text()
      case "%ARCH%":
         scan.Scan()
         desc.arch = scan.Text()
      case "%PROVIDES%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" { break }
            db.provides[baseName(line, ">=")] = name
         }
      case "%DEPENDS%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" { break }
            desc.depends = append(desc.depends, baseName(line, ">="))
         }
         db.name[name] = desc
      }
   }
   return nil
}

func (db database) sync(txt string) error {
   open, e := os.Open(txt)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      name := scan.Text()
      println(db.name[name].filename)
      // download file and extract
   }
   return nil
}

type description struct {
   filename string
   arch string
   depends []string
}

type install struct {
   source string
   cache string
   dest string
}

func newInstall(source url.URL, base ...string) (install, error) {
   cache, e := os.UserCacheDir()
   if e != nil {
      return install{}, e
   }
   dest := filepath.VolumeName(cache) + string(os.PathSeparator)
   for _, each := range base {
      cache = filepath.Join(cache, each)
      dest = filepath.Join(dest, each)
   }
   src := source.String()
   cache = filepath.Join(cache, filepath.Base(src))
   return install{src, cache, dest}
}
