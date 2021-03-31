package main

import (
   "bufio"
   "bytes"
   "net/url"
   "os"
   "strings"
)

var mirror = url.URL{Scheme: "http", Host: "repo.msys2.org"}

func baseName(s, chars string) string {
   n := strings.IndexAny(s, chars)
   if n == -1 { return s }
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
      buf = bufio.NewScanner(bytes.NewReader(file))
   )
   for buf.Scan() {
      switch buf.Text() {
      case "%FILENAME%":
         buf.Scan()
         desc.filename = buf.Text()
      case "%NAME%":
         buf.Scan()
         name = buf.Text()
      case "%ARCH%":
         buf.Scan()
         desc.arch = buf.Text()
      case "%PROVIDES%":
         for buf.Scan() {
            line := buf.Text()
            if line == "" { break }
            db.provides[baseName(line, ">=")] = name
         }
      case "%DEPENDS%":
         for buf.Scan() {
            line := buf.Text()
            if line == "" { break }
            desc.depends = append(desc.depends, baseName(line, ">="))
         }
         db.name[name] = desc
      }
   }
   return nil
}

type description struct {
   filename string
   arch string
   depends []string
}

func (db database) sync(name string) error {
   file, e := os.Open(name)
   if e != nil { return e }
   defer file.Close()
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      println(db.name[buf.Text()].filename)
      // FIXME download file and extract
   }
   return nil
}
