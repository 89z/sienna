package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "sienna"
)

var toml_sep = []byte{'+', '+', '+', '\n'}

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func tomlDecode(y []byte) (sienna.Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}
