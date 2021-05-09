package main

import (
   "bufio"
   "bytes"
   "fmt"
   "github.com/89z/rosso"
   "os"
   "path/filepath"
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
   file, err := os.Open(name)
   if err != nil { return err }
   defer file.Close()
   buf := bufio.NewScanner(file)
   // cache
   cache, err := os.UserCacheDir()
   if err != nil { return err }
   cache = filepath.Join(cache, "sienna", "msys2")
   for buf.Scan() {
      text := buf.Text()
      var base string
      if strings.Contains(text, ".pkg.tar.") {
         base = text
      } else {
         desc, ok := db.name[text]
         if ! ok {
            fmt.Printf("%q not valid\n", text)
            continue
         }
         base = desc.filename
      }
      // get
      get := mirror + variant(base) + base
      // create
      create := filepath.Join(cache, base)
      // copy
      err := rosso.Copy(get, create)
      if os.IsExist(err) {
         fmt.Println("Exist", base)
      } else if err != nil {
         return err
      }
      var tar rosso.Archive
      switch filepath.Ext(base) {
      case ".xz":
         fmt.Println("Xz", base)
         tar.Xz(create, `C:\sienna\msys2`)
      case ".zst":
         fmt.Println("Zst", base)
         tar.Zst(create, `C:\sienna\msys2`)
      }
   }
   return nil
}

type description struct {
   filename string
   depends []string
}

func main() {
   if len(os.Args) != 3 {
      fmt.Println(`install-msys2 query git
install-msys2 sync git.txt`)
      os.Exit(1)
   }
   data := newDatabase()
   // cache
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "msys2")
   for _, db := range []string{
      "/mingw/ucrt64/ucrt64.db",
      "/mingw/x86_64/mingw64.db",
      "/msys/x86_64/msys.db",
   } {
      // create
      create := filepath.Join(cache, db)
      err := rosso.Copy(mirror + db, create)
      if os.IsExist(err) {
         fmt.Println("Exist", db)
      } else if err != nil {
         panic(err)
      }
      fs, err := rosso.TarGzMemory(create)
      if err != nil {
         panic(err)
      }
      for _, file := range fs {
         data.scan(file.Data)
      }
   }
   target := os.Args[2]
   switch os.Args[1] {
   case "query": data.query(target)
   case "sync": data.sync(target)
   }
}
