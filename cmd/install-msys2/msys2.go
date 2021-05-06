package main

import (
   "bufio"
   "bytes"
   "fmt"
   "github.com/89z/rosso"
   "os"
   "path"
   "strings"
)

const mirror = "http://repo.msys2.org"

func variant(s string) string {
   switch {
   case strings.HasPrefix(s, "mingw-w64-ucrt-x86_64-"):
      return "/mingw/ucrt64/"
   case strings.HasPrefix(s, "mingw-w64-x86_64-"):
      return "/mingw/x86_64/"
   default:
      return "/msys/x86_64/"
   }
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

func baseName(s, chars string) string {
   n := strings.IndexAny(s, chars)
   switch n {
   case -1: return s
   default: return s[:n]
   }
}

func (db database) sync(name string) error {
   file, e := os.Open(name)
   if e != nil { return e }
   defer file.Close()
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      text := buf.Text()
      var filename string
      if strings.Contains(text, ".pkg.tar.") {
         filename = text
      } else {
         desc, ok := db.name[text]
         if ! ok {
            fmt.Printf("%q not valid\n", text)
            continue
         }
         filename = desc.filename
      }
      inst := rosso.NewInstall("sienna/msys2", filename)
      inst.SetCache()
      dir := variant(filename)
      _, e = rosso.Copy(mirror + dir + filename, inst.Cache)
      if os.IsExist(e) {
         rosso.LogInfo("Exist", filename)
      } else if e != nil {
         return e
      }
      var arc rosso.Archive
      switch path.Ext(filename) {
      case ".xz":
         rosso.LogInfo("Xz", filename)
         arc.Xz(inst.Cache, inst.Dest)
      case ".zst":
         rosso.LogInfo("Zst", filename)
         e = arc.Zst(inst.Cache, inst.Dest)
         if e != nil { return e }
      }
   }
   return nil
}

type description struct {
   filename string
   depends []string
}
