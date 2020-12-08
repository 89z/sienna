package main

import (
   "encoding/json"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Map map[string]interface{}

func (m Map) M(s string) Map {
   return m[s].(map[string]interface{})
}

func Encode(m map[string]interface{}, s string) error {
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   return toml.
   NewEncoder(o).
   ArraysWithOneElementPerLine(true).
   QuoteMapKeys(true).
   Encode(m)
}

func main() {
   if len(os.Args) != 2 {
      println("toml <json file>")
      os.Exit(1)
   }

   path_s := os.Args[1]
   open_o, e := os.Open(path_s)
   if e != nil {
      log.Fatal(e)
   }

   artist_m := Map{}
   json.NewDecoder(open_o).Decode(&artist_m)
   for s := range artist_m.M("Cocteau Twins") {
      println(s)
   }
}
