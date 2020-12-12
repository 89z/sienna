package manager

import (
   "bufio"
   "fmt"
   "github.com/mholt/archiver/v3"
   "io/ioutil"
   "msys2/net"
   "os"
   "path"
   "strings"
)

type Manager struct {
   Cache string
   Packages []os.FileInfo
}

func NewManager() (Manager, error) {
   cache_s, e := os.UserCacheDir()
   if e != nil {
      return Manager{}, e
   }
   msys_s := path.Join(cache_s, "Msys2")
   dir_a, e := ioutil.ReadDir(msys_s)
   if e != nil {
      return Manager{}, e
   }
   db_a := []string{"mingw64.db.tar.gz", "msys.db.tar.gz"}
   for n := range db_a {
      file_s := db_a[n]
      real_s := path.Join(msys_s, file_s)
      if IsFile(real_s) {
         continue
      }
      url_s := GetRepo(file_s) + file_s
      e = net.Copy(url_s, real_s)
      if e != nil {
         return Manager{}, e
      }
      e = Unarchive(real_s, msys_s)
      if e != nil {
         return Manager{}, e
      }
   }
   return Manager{msys_s, dir_a}, nil
}

func (o Manager) GetName(pack_s string) (string, error) {
   for n := range o.Packages {
      dir_s := o.Packages[n].Name()
      if strings.HasPrefix(dir_s, pack_s + "-") {
         return dir_s, nil
      }
   }
   return "", fmt.Errorf(pack_s)
}

func (o Manager) GetValue(pack_s, key_s string) ([]string, error) {
   a := []string{}
   name_s, e := o.GetName(pack_s)
   if e != nil {
      return a, e
   }
   real_s := path.Join(o.Cache, name_s, "desc")
   open_o, e := os.Open(real_s)
   if e != nil {
      return a, e
   }
   scan_o := bufio.NewScanner(open_o)
   dep_b := false
   for scan_o.Scan() {
      line_s := scan_o.Text()
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
      a = append(a, BaseName(line_s, "=>"))
   }
   return a, nil
}

func (o Manager) Resolve(pack_s string) (map[string]bool, error) {
   pack_m := map[string]bool{}
   for pack_a := []string{pack_s}; len(pack_a) > 0; pack_a = pack_a[1:] {
      pack_s := pack_a[0]
      dep_a, e := o.GetValue(pack_s, "%DEPENDS%")
      if e != nil {
         return pack_m, e
      }
      pack_m[pack_s] = true
      pack_a = append(pack_a, dep_a...)
   }
   return pack_m, nil
}

func (o Manager) Sync(tar_s string) error {
   open_o, e := os.Open(tar_s)
   if e != nil {
      return e
   }
   scan_o := bufio.NewScanner(open_o)
   for scan_o.Scan() {
      pack_s := scan_o.Text()
      val_a, e := o.GetValue(pack_s, "%FILENAME%")
      if e != nil {
         return e
      }
      file_s := val_a[0]
      real_s := path.Join(o.Cache, file_s)
      if ! IsFile(real_s) {
         url_s := GetRepo(file_s) + file_s
         e := net.Copy(url_s, real_s)
         if e != nil {
            return e
         }
      }
      e = Unarchive(real_s, `C:\msys2`)
      if e != nil {
         return e
      }
   }
   return nil
}

func BaseName(s, char_s string) string {
   n := strings.IndexAny(s, char_s)
   if n == -1 {
      return s
   }
   return s[:n]
}

func GetRepo(file_s string) string {
   switch file_s[:17] {
   case "mingw-w64-x86_64-", "mingw64.db.tar.gz":
      return "http://repo.msys2.org/mingw/x86_64/"
   default:
      return "http://repo.msys2.org/msys/x86_64/"
   }
}

func IsFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

func Unarchive(in_path, out_path string) error {
   tar_o := &archiver.Tar{OverwriteExisting: true}
   in_file := path.Base(in_path)
   fmt.Println("EXTRACT", in_file)
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
