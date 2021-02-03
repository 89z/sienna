package main

import (
   "bufio"
   "errors"
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "io/ioutil"
   "os"
   "path"
   "strings"
)

func baseName(s, char_s string) string {
   n := strings.IndexAny(s, char_s)
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

type manager struct {
   cache string
   packages []os.FileInfo
}

func newManager() (m manager, e error) {
   cache, e := x.GetCache("msys2")
   if e != nil {
      return
   }
   dir, e := ioutil.ReadDir(cache)
   if e != nil {
      return
   }
   for _, file := range []string{"mingw64.db.tar.gz", "msys.db.tar.gz"} {
      abs := path.Join(cache, file)
      if x.IsFile(abs) {
         continue
      }
      _, e = x.Copy(
         getRepo(file) + file, abs,
      )
      if e != nil {
         return
      }
      e = unarchive(abs, cache)
      if e != nil {
         return
      }
   }
   return manager{cache, dir}, nil
}

func (m manager) getName(pack string) (string, error) {
   for n := range m.packages {
      dir_s := m.packages[n].Name()
      if strings.HasPrefix(dir_s, pack + "-") {
         return dir_s, nil
      }
   }
   return "", errors.New(pack)
}

func (m manager) getValue(pack, key_s string) (a []string, e error) {
   name, e := m.getName(pack)
   if e != nil {
      return
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
      if line == key_s {
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

func (m manager) resolve(pack string) (map[string]bool, error) {
   set := map[string]bool{}
   for packs := []string{pack}; len(packs) > 0; packs = packs[1:] {
      pack := packs[0]
      deps, e := m.getValue(pack, "%DEPENDS%")
      if e != nil {
         return set, e
      }
      set[pack] = true
      packs = append(packs, deps...)
   }
   return set, nil
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
   man, e := newManager()
   x.Check(e)
   if os.Args[1] == "deps" {
      deps, e := man.resolve(target)
      x.Check(e)
      for dep := range deps {
         println(dep)
      }
   } else {
      e := man.sync(target)
      x.Check(e)
   }
}
