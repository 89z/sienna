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

func variant(s string) string {
   switch {
   case strings.HasPrefix(s, "mingw-w64-ucrt-x86_64-"):
      return "/mingw/ucrt64/"
   case strings.HasPrefix(s, "mingw-w64-x86_64-"):
      return "/mingw/x86_64/"
   default:
      return "/msys/x86_64/"
   }
}

func (db database) sync(name string) error {
   file, e := os.Open(name)
   if e != nil { return e }
   defer file.Close()
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      text := buf.Text()
      desc, ok := db.name[text]
      if ! ok {
         fmt.Printf("%q not valid\n", text)
         continue
      }
      inst := x.NewInstall("sienna/msys2", desc.filename)
      inst.SetCache()
      dir := variant(desc.filename)
      _, e = x.Copy(mirror + dir + desc.filename, inst.Cache)
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
