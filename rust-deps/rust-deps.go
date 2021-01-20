package main

import (
   "github.com/89z/sienna"
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

type Map = sienna.Map
