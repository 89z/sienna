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
      url := getRepo(file) + file
      _, e = x.Copy(url, abs)
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

func (m manager) getValue(pack, key_s string) ([]string, error) {
   a := []string{}
   name, e := m.getName(pack)
   if e != nil {
      return a, e
   }
   abs := path.Join(m.cache, name, "desc")
   open, e := os.Open(abs)
   if e != nil {
      return a, e
   }
   scan := bufio.NewScanner(open)
   dep_b := false
   for scan.Scan() {
      line_s := scan.Text()
      // STATE 2
      if line_s == key_s {
         dep_b = true
         continue
      }
      // STATE 1
      if ! dep_b {
         continue
      }
      // STATE 4
      if line_s == "" {
         break
      }
      // STATE 3
      a = append(a, baseName(line_s, "=>"))
   }
   return a, nil
}

func (m manager) resolve(pack string) (map[string]bool, error) {
   pack_m := map[string]bool{}
   for pack_a := []string{pack}; len(pack_a) > 0; pack_a = pack_a[1:] {
      pack := pack_a[0]
      dep_a, e := m.getValue(pack, "%DEPENDS%")
      if e != nil {
         return pack_m, e
      }
      pack_m[pack] = true
      pack_a = append(pack_a, dep_a...)
   }
   return pack_m, nil
}

func (m manager) sync(tar_s string) error {
   open, e := os.Open(tar_s)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      pack := scan.Text()
      val_a, e := m.getValue(pack, "%FILENAME%")
      if e != nil {
         return e
      }
      file_s := val_a[0]
      abs := path.Join(m.cache, file_s)
      if ! x.IsFile(abs) {
         url := getRepo(file_s) + file_s
         _, e := x.Copy(url, abs)
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
