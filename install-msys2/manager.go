package main

import (
   "bufio"
   "bytes"
   "fmt"
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "net/url"
   "os"
   "path"
   "strings"
)

var mirror = url.URL{Scheme: "http", Host: "repo.msys2.org"}

func baseName(s, chars string) string {
   n := strings.IndexAny(s, chars)
   if n == -1 { return s }
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

func (db database) query(target string) {
   done := map[string]bool{target: true}
   for todo := []string{target}; len(todo) > 0; todo = todo[1:] {
      do := todo[0]
      for _, dep := range db.name[do].depends {
         provide, ok := db.provides[dep]
         if ok { dep = provide }
         if ! done[dep] {
            todo = append(todo, dep)
            done[dep] = true
         }
      }
      fmt.Println(do)
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
      }
   }
   db.name[name] = desc
   return nil
}

func (db database) sync(name string) error {
   file, e := os.Open(name)
   if e != nil { return e }
   defer file.Close()
   buf := bufio.NewScanner(file)
   repos := map[bool]string{true: "mingw", false: "msys"}
   for buf.Scan() {
      desc, ok := db.name[buf.Text()]
      if ! ok { continue }
      repo := repos[strings.HasPrefix(desc.filename, "mingw-w64-x86_64-")]
      mirror.Path = repo + "/x86_64/" + desc.filename
      inst := x.NewInstall("sienna/msys2", desc.filename)
      inst.SetCache()
      _, e = x.Copy(mirror.String(), inst.Cache)
      if os.IsExist(e) {
         x.LogInfo("Exist", desc.filename)
      } else if e != nil {
         return e
      }
      var arc extract.Archive
      switch path.Ext(desc.filename) {
      case ".xz":
         x.LogInfo("Xz", desc.filename)
         arc.Xz(inst.Cache, inst.Dest)
      case ".zst":
         x.LogInfo("Zst", desc.filename)
         arc.Zst(inst.Cache, inst.Dest)
      }
   }
   return nil
}

type description struct {
   filename string
   depends []string
}
