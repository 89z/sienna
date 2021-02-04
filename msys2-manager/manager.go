package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "io/ioutil"
   "os"
   "path/filepath"
   "strings"
)

type manager struct {
   x.Install
}

func (m manager) getValue(pack, key string) (val []string, e error) {
   var name string
   packages, e := ioutil.ReadDir(m.Cache)
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
   if name == "" {
      return nil, fmt.Errorf("%v %v", pack, key)
   }
   open, e := os.Open(
      filepath.Join(m.Cache, name, "desc"),
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
      base := baseName(line, "=>")
      if base == "sh" {
         return nil, fmt.Errorf("%v %v %v", name, key, line)
      }
      val = append(val, base)
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
      values, e := m.getValue(
         scan.Text(), "%FILENAME%",
      )
      if e != nil {
         return e
      }
      file := values[0]
      archive := filepath.Join(m.Cache, file)
      if ! x.IsFile(archive) {
         _, e := x.Copy(
            getRepo(file) + file, archive,
         )
         if e != nil {
            return e
         }
      }
      e = unarchive(archive, m.Dest)
      if e != nil {
         return e
      }
   }
   return nil
}
