package main

import (
   "strings"
   "testing"
)

var tests = []struct{
   in string
   out string
}{
   {
      "bluewednesday/murmuration-feat-shopan",
      `",2019,"s/593405433/000507498393-dd22sy","Murmuration (feat. Shopan)"]`,
   }, {
      "four-tet/burial-four-tet-nova",
      `",2012,"s/38720262/000014893963-91gp52","Burial + Four Tet - Nova"]`,
   },
}

func TestInsert(t *testing.T) {
   for _, each := range tests {
      out, e := insert("https://soundcloud.com/" + each.in)
      if e != nil {
         t.Error(e)
      }
      if ! strings.HasSuffix(out, each.out) {
         t.Error(out)
      }
   }
}
