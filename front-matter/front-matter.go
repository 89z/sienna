package main

import (
   "bytes"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
   "os"
   "path"
   "sienna"
)

var toml_sep = []byte{'+', '+', '+', '\n'}

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}
