package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "os"
   "path/filepath"
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
   base := filepath.Base(in)
   println(x.ColorCyan("Extract"), base)
   switch filepath.Ext(base) {
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
      archive := filepath.Join(install.Cache, each)
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
