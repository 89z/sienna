package main

import (
   "bufio"
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "os"
   "path"
   "strings"
)

func (db database) sync(name string) error {
   file, e := os.Open(name)
   if e != nil { return e }
   defer file.Close()
   buf := bufio.NewScanner(file)
   repos := map[bool]string{true: "mingw", false: "msys"}
   for buf.Scan() {
      text := buf.Text()
      desc, ok := db.name[text]
      if ! ok {
         provide, ok := db.provides[text]
         if ! ok { continue }
         desc = db.name[provide]
      }
      repo := repos[strings.HasPrefix(desc.filename, "mingw-w64-x86_64-")]
      mirror.Path = repo + "/x86_64/" + desc.filename
      inst := x.NewInstall("sienna/msys2", desc.filename)
      inst.SetCache()
      _, e = x.Copy(mirror.String(), inst.Cache)
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
         arc.Zst(inst.Cache, inst.Dest)
      }
   }
   return nil
}
