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

func unarchive(in_path, out_path string) error {
   tar_o := &archiver.Tar{OverwriteExisting: true}
   in_file := path.Base(in_path)
   println("EXTRACT", in_file)
   switch path.Ext(in_file) {
   case ".zst":
      zstd_o := archiver.TarZstd{Tar: tar_o}
      return zstd_o.Unarchive(in_path, out_path)
   case ".xz":
      xz_o := archiver.TarXz{Tar: tar_o}
      return xz_o.Unarchive(in_path, out_path)
   default:
      gz_o := archiver.TarGz{Tar: tar_o}
      return gz_o.Unarchive(in_path, out_path)
   }
}

type manager struct {
   cache string
   packages []os.FileInfo
}

func newManager() (manager, error) {
   cache, e := os.UserCacheDir()
   if e != nil {
      return manager{}, e
   }
   msys_s := path.Join(cache, "Msys2")
   dir_a, e := ioutil.ReadDir(msys_s)
   if e != nil {
      return manager{}, e
   }
   db_a := []string{"mingw64.db.tar.gz", "msys.db.tar.gz"}
   for n := range db_a {
      file_s := db_a[n]
      real_s := path.Join(msys_s, file_s)
      if x.IsFile(real_s) {
         continue
      }
      url := getRepo(file_s) + file_s
      _, e = x.HttpCopy(url, real_s)
      if e != nil {
         return manager{}, e
      }
      e = unarchive(real_s, msys_s)
      if e != nil {
         return manager{}, e
      }
   }
   return manager{msys_s, dir_a}, nil
}

func (o manager) getName(pack string) (string, error) {
   for n := range o.packages {
      dir_s := o.packages[n].Name()
      if strings.HasPrefix(dir_s, pack + "-") {
         return dir_s, nil
      }
   }
   return "", errors.New(pack)
}

func (o manager) getValue(pack, key_s string) ([]string, error) {
   a := []string{}
   name, e := o.getName(pack)
   if e != nil {
      return a, e
   }
   real_s := path.Join(o.cache, name, "desc")
   open, e := os.Open(real_s)
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

func (o manager) resolve(pack string) (map[string]bool, error) {
   pack_m := map[string]bool{}
   for pack_a := []string{pack}; len(pack_a) > 0; pack_a = pack_a[1:] {
      pack := pack_a[0]
      dep_a, e := o.getValue(pack, "%DEPENDS%")
      if e != nil {
         return pack_m, e
      }
      pack_m[pack] = true
      pack_a = append(pack_a, dep_a...)
   }
   return pack_m, nil
}

func (o manager) sync(tar_s string) error {
   open, e := os.Open(tar_s)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      pack := scan.Text()
      val_a, e := o.getValue(pack, "%FILENAME%")
      if e != nil {
         return e
      }
      file_s := val_a[0]
      real_s := path.Join(o.cache, file_s)
      if ! x.IsFile(real_s) {
         url := getRepo(file_s) + file_s
         _, e := x.HttpCopy(url, real_s)
         if e != nil {
            return e
         }
      }
      e = unarchive(real_s, `C:\msys2`)
      if e != nil {
         return e
      }
   }
   return nil
}
