package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "os"
   "path"
)

const repo = "http://repo.msys2.org"

type desc struct {
   name string
   filename string
   provides []string
   depends []string
}

func (m manager) getValue(pack, key string) ([]string, error) {
   packages, e := ioutil.ReadDir(m.Cache)
   if e != nil {
      return nil, e
   }
   var (
      dep bool
      name string
      val []string
   )
   for _, each := range packages {
      dir := each.Name()
      if strings.HasPrefix(dir, pack + "-") {
         name = dir
         break
      }
   }
   if name == "" {
      return nil, fmt.Errorf("%v %v", pack, key)
   }
   open, e := os.Open(
      path.Join(m.Cache, name, "desc"),
   )
   if e != nil {
      return nil, e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      line := scan.Text()
      // STATE 2
      if line == key {
         dep = true
         continue
      }
      // STATE 1
      if ! dep {
         continue
      }
      // STATE 4
      if line == "" {
         break
      }
      // STATE 3
      base := baseName(line, "=>")
      val = append(val, base)
   }
   return val, nil
}

func main() {
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
      archive := cache + each
      _, e = x.Copy(repo + each, archive)
      if e != nil {
         log.Fatal(e)
      }
      e = tar.Gz(
         archive, path.Dir(archive),
      )
      if e != nil {
         log.Fatal(e)
      }
   }
}
