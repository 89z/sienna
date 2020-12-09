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
   if len(os.Args) != 3 {
      println("toml <json file> <toml file>")
      os.Exit(1)
   }
   in_s := os.Args[1]
   out_s := os.Args[2]
   open_o, e := os.Open(in_s)
   if e != nil {
      log.Fatal(e)
   }
   artist_m := Map{}
   json.NewDecoder(open_o).Decode(&artist_m)
   e = Encode(artist_m, out_s)
   if e != nil {
      log.Fatal(e)
   }
}
