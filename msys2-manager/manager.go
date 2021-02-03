package main

import (
   "bufio"
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "io/ioutil"
   "os"
   "path"
   "strings"
)

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

func unarchive(in, out string) error {
   tar := &archiver.Tar{OverwriteExisting: true}
   base := path.Base(in)
   println("EXTRACT", base)
   switch path.Ext(base) {
   case ".zst":
      zs := archiver.TarZstd{Tar: tar}
      return zs.Unarchive(in, out)
   case ".xz":
      xz := archiver.TarXz{Tar: tar}
      return xz.Unarchive(in, out)
   default:
      gz := archiver.TarGz{Tar: tar}
      return gz.Unarchive(in, out)
   }
}


func (m manager) getValue(pack, key string) (a []string, e error) {
   var name string
   packages, e := ioutil.ReadDir(m.cache)
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
   abs := path.Join(m.cache, name, "desc")
   open, e := os.Open(abs)
   if e != nil {
      return
   }
   dep := false
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
      a = append(a, baseName(line, "=>"))
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
      pack := scan.Text()
      values, e := m.getValue(pack, "%FILENAME%")
      if e != nil {
         return e
      }
      file := values[0]
      abs := path.Join(m.cache, file)
      if ! x.IsFile(abs) {
         _, e := x.Copy(
            getRepo(file) + file, abs,
         )
         if e != nil {
            return e
         }
      }
      e = unarchive(abs, `C:\msys2`)
      if e != nil {
         return e
      }
   }
   return nil
}


func main() {
   if len(os.Args) != 3 {
      println(`synopsis:
   msys2 <operation> <target>

examples:
   msys2 deps mingw-w64-x86_64-libgit2
   msys2 sync git.txt`)
      os.Exit(1)
   }
   target := os.Args[2]
   install, e := x.NewInstall("msys2")
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
      e = unarcive(archive, install.Cache)
      x.Check(e)
   }
   if os.Args[1] == "sync" {
      e = sync(install, target)
      x.Check(e)
      return
   }
   // var packSet = map[string]bool{}
   for packs := []string{target}; len(packs) > 0; packs = packs[1:] {
      target := packs[0]
      // FIXME
      deps, e := m.getValue(target, "%DEPENDS%")
      if e != nil {
         return packSet, e
      }
      packSet[target] = true
      packs = append(packs, deps...)
   }
   deps, e := resolve(install, target)
   x.Check(e)
   for dep := range deps {
      println(dep)
   }
}
