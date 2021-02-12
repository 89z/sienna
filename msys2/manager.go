package main

import (
   "bufio"
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
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

func getDesc(file, repo, variant string) (description, error) {
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

func main() {
   if len(os.Args) != 3 {
      println(`msys2 query mingw-w64-x86_64-gcc
msys2 sync gcc.txt`)
      os.Exit(1)
   }
   cache, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }
   cache = path.Join(cache, "sienna")
   var tar extract.Tar
   for _, each := range []string{
      "/mingw/x86_64/mingw64.db.tar.gz",
      "/msys/x86_64/msys.db.tar.gz",
   } {
      archive := path.Join(cache, each)
      _, e = x.Copy(mirror + each, archive)
      if os.IsExist(e) {
         continue
      } else if e != nil {
         log.Fatal(e)
      }
      println(x.ColorCyan("Extract"), each)
      e = tar.Gz(
         archive, path.Dir(archive),
      )
      if e != nil {
         log.Fatal(e)
      }
   }
}
