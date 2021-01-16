package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
   "sienna"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func tomlDecode(s string) (Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func tomlEncode(s string, m Map) error {
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   defer o.Close()
   return toml.NewEncoder(o).Encode(m)
}

type Map = sienna.Map
