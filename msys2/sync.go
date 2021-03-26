package main

import (
   "bufio"
   "os"
)

func (db database) sync(txt string) error {
   open, e := os.Open(txt)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      name := scan.Text()
      println(db.name[name].filename)
      // download file and extract
   }
   return nil
}
