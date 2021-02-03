package main

import (
   "bufio"
   "github.com/89z/x"
   "io/ioutil"
   "os"
   "path"
   "strings"
)

func getValue(install x.Install, pack, key string) (val []string, e error) {
   var name string
   packages, e := ioutil.ReadDir(install.Cache)
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
   open, e := os.Open(
      path.Join(install.Cache, name, "desc"),
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
      val = append(val, baseName(line, "=>"))
   }
   return
}

func sync(install x.Install, tar string) error {
   open, e := os.Open(tar)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      values, e := getValue(
         install, scan.Text(), "%FILENAME%",
      )
      if e != nil {
         return e
      }
      file := values[0]
      archive := path.Join(install.Cache, file)
      if ! x.IsFile(archive) {
         _, e := x.Copy(
            getRepo(file) + file, archive,
         )
         if e != nil {
            return e
         }
      }
      e = unarchive(archive, install.Dest)
      if e != nil {
         return e
      }
   }
   return nil
}
