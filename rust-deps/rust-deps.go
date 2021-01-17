package main

import (
   "log"
   "os"
   "sienna"
)

type Map = sienna.Map

var (
   dep int
   prev string
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}
