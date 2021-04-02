package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "os"
   "path"
   "strings"
)

/*
bad
http://repo.msys2.org/msys/x86_64/mingw-w64-ucrt-x86_64-gcc-10.2.0-9-any.pkg.tar.zst

good
http://repo.msys2.org/mingw/ucrt64/mingw-w64-ucrt-x86_64-gcc-10.2.0-9-any.pkg.tar.zst
*/

func (db database) sync(name string) error {
   file, e := os.Open(name)
   if e != nil { return e }
   defer file.Close()
   buf := bufio.NewScanner(file)
   repos := map[bool]string{true: "/mingw", false: "/msys"}
   for buf.Scan() {
      text := buf.Text()
      desc, ok := db.name[text]
      if ! ok {
         fmt.Printf("%q not valid\n", text)
         continue
      }
      inst := x.NewInstall("sienna/msys2", desc.filename)
      inst.SetCache()
      repo := repos[strings.HasPrefix(desc.filename, "mingw-w64-x86_64-")]
      _, e = x.Copy(mirror + repo + "/x86_64/" + desc.filename, inst.Cache)
      if os.IsExist(e) {
         x.LogInfo("Exist", desc.filename)
      } else if e != nil {
         return e
      }
      var arc extract.Archive
      switch path.Ext(desc.filename) {
      case ".xz":
         x.LogInfo("Xz", desc.filename)
         arc.Xz(inst.Cache, inst.Dest)
      case ".zst":
         x.LogInfo("Zst", desc.filename)
         e = arc.Zst(inst.Cache, inst.Dest)
         if e != nil { return e }
      }
   }
   return nil
}
