package main

import (
   "github.com/89z/sienna"
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "log"
   "os"
)

var (
   dep int
   prev string
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}
