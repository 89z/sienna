package main

import (
   "archive/tar"
   "bufio"
   "bytes"
   "compress/gzip"
   "fmt"
   "github.com/89z/sienna"
   "io"
   "net/http"
   "os"
   "path/filepath"
   "strings"
   "testing/fstest"
)

const mirror = "http://repo.msys2.org"

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func tarGzMemory(source string) (fstest.MapFS, error) {
   file, err := os.Open(source)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   gzRead, err := gzip.NewReader(file)
   if err != nil {
      return nil, err
   }
   tarRead := tar.NewReader(gzRead)
   files := make(fstest.MapFS)
   for {
      cur, err := tarRead.Next()
      if err == io.EOF {
         break
      } else if err != nil {
         return nil, err
      }
      if cur.Typeflag != tar.TypeReg {
         continue
      }
      data, err := io.ReadAll(tarRead)
      if err != nil {
         return nil, err
      }
      files[cur.Name] = &fstest.MapFile{Data: data}
   }
   return files, nil
}

func variant(s string) string {
   if strings.HasPrefix(s, "mingw-w64-ucrt-x86_64-") {
      return "/mingw/ucrt64/"
   }
   if strings.HasPrefix(s, "mingw-w64-x86_64-") {
      return "/mingw/x86_64/"
   }
   return "/msys/x86_64/"
}

type database struct {
   // %NAME% -> %FILENAME%, %DEPENDS%
   name map[string]description
   // %PROVIDES% -> %NAME%
   provides map[string]string
}

func newDatabase() database {
   return database{
      make(map[string]description), make(map[string]string),
   }
}

func (db database) query(target string) {
   done := map[string]bool{target: true}
   for todo := []string{target}; len(todo) > 0; todo = todo[1:] {
      do := todo[0]
      for _, dep := range db.name[do].depends {
         provide, ok := db.provides[dep]
         if ok {
            dep = provide
         }
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
            if line == "" {
               break
            }
            db.provides[baseName(line, ">=")] = name
         }
      case "%DEPENDS%":
         for buf.Scan() {
            line := buf.Text()
            if line == "" {
               break
            }
            desc.depends = append(desc.depends, baseName(line, ">="))
         }
      }
   }
   db.name[name] = desc
   return nil
}

func baseName(s, chars string) string {
   if n := strings.IndexAny(s, chars); n >= 0 {
      return s[:n]
   }
   return s
}

func (db database) sync(name string) error {
   file, err := os.Open(name)
   if err != nil {
      return err
   }
   defer file.Close()
   buf := bufio.NewScanner(file)
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
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
      create := filepath.Join(cache, base)
      _, err := os.Stat(create)
      if err != nil {
         r, err := http.Get(mirror + variant(base) + base)
         if err != nil {
            return err
         }
         defer r.Body.Close()
         f, err := os.Create(create)
         if err != nil {
            return err
         }
         defer f.Close()
         f.ReadFrom(r.Body)
      } else {
         fmt.Println(invert, "Exist", reset, base)
      }
      var tar sienna.Archive
      switch filepath.Ext(base) {
      case ".xz":
         fmt.Println(invert, "Xz", reset, base)
         tar.Xz(create, `C:\sienna\msys2`)
      case ".zst":
         fmt.Println(invert, "Zst", reset, base)
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
      return
   }
   data := newDatabase()
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
      create := filepath.Join(cache, db)
      _, err := os.Stat(create)
      if err != nil {
         r, err := http.Get(mirror + db)
         if err != nil {
            panic(err)
         }
         defer r.Body.Close()
         f, err := os.Create(create)
         if err != nil {
            panic(err)
         }
         defer f.Close()
         f.ReadFrom(r.Body)
      } else {
         fmt.Println(invert, "Exist", reset, db)
      }
      fs, err := tarGzMemory(create)
      if err != nil {
         panic(err)
      }
      for _, file := range fs {
         data.scan(file.Data)
      }
   }
   target := os.Args[2]
   if os.Args[1] == "sync" {
      data.sync(target)
      return
   }
   data.query(target)
}
