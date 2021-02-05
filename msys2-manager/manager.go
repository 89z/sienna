package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "io/ioutil"
   "os"
   "path"
   "strings"
)

type manager struct {
   x.Install
}

func (m manager) getValue(pack, key string) (val []string, e error) {
   var name string
   packages, e := ioutil.ReadDir(m.Cache)
   if e != nil {
      return
   }
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
      return
   }
   var dep bool
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
      if base == "sh" {
         return nil, fmt.Errorf("%v %v %v", name, key, line)
      }
      val = append(val, base)
   }
   return
}

func (m manager) sync(tar string) error {
   open, e := os.Open(tar)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      values, e := m.getValue(
         scan.Text(), "%FILENAME%",
      )
      if e != nil {
         return e
      }
      file := values[0]
      archive := path.Join(m.Cache, file)
      if ! x.IsFile(archive) {
         _, e := x.Copy(
            getRepo(file) + file, archive,
         )
         if e != nil {
            return e
         }
      }
      e = unarchive(archive, m.Dest)
      if e != nil {
         return e
      }
   }
   return nil
}



func baseName(s, char string) string {
   n := strings.IndexAny(s, char)
   if n == -1 {
      return s
   }
   return s[:n]
}

func getRepo(s string) string {
   if s == "mingw64.db.tar.gz" || strings.HasPrefix(s, "mingw-w64-x86_64-") {
      return "http://repo.msys2.org/mingw/x86_64/"
   }
   return "http://repo.msys2.org/msys/x86_64/"
}

func unarchive(source, dest string) error {
   var tar extract.Tar
   switch path.Ext(source) {
   case ".zst":
      return tar.Zst(source, dest)
   case ".xz":
      return tar.Xz(source, dest)
   default:
      return tar.Gz(source, dest)
   }
}

func main() {
   if len(os.Args) != 3 {
      println(`synopsis:
   msys2-manager <operation> <target>

examples:
   msys2-manager deps mingw-w64-x86_64-libgit2
   msys2-manager sync git.txt`)
      os.Exit(1)
   }
   target := os.Args[2]
   install, e := x.NewInstall("msys64")
   x.Check(e)
   for _, each := range []string{"mingw64.db.tar.gz", "msys.db.tar.gz"} {
      archive := path.Join(install.Cache, each)
      if x.IsFile(archive) {
         continue
      }
      _, e = x.Copy(
         getRepo(each) + each, archive,
      )
      x.Check(e)
      e = unarchive(archive, install.Cache)
      x.Check(e)
   }
   man := manager{install}
   if os.Args[1] == "sync" {
      e = man.sync(target)
      x.Check(e)
      return
   }
   var packSet = map[string]bool{}
   for packs := []string{target}; len(packs) > 0; packs = packs[1:] {
      target := packs[0]
      deps, e := man.getValue(target, "%DEPENDS%")
      x.Check(e)
      packs = append(packs, deps...)
      if packSet[target] {
         continue
      }
      println(target)
      packSet[target] = true
   }
}
