package main

import (
   "bytes"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
   "os"
   "sienna"
)

var toml_sep = []byte{'+', '+', '+', '\n'}

func isFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

func tomlDecode(y []byte) (sienna.Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   os.Chdir(`D:\Git\autumn\content`)
   dir_a, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for _, entry := range dir_a {
      index_s := entry.Name() + `\_index.md`
      index_y, e := ioutil.ReadFile(index_s)
      if e != nil {
         log.Fatal(e)
      }
      toml_y := bytes.SplitN(index_y, toml_sep, 3)[1]
      toml_m, e := tomlDecode(toml_y)
      if e != nil {
         log.Fatal(e)
      }
      if toml_m["_build"] != nil {
         continue
      }
      example_s := `D:\Git\` + toml_m.A("example").S(0)
      if ! isFile(example_s) {
         println(index_s)
         continue
      }
      example_y, e := ioutil.ReadFile(example_s)
      if e != nil {
         log.Fatal(e)
      }
      substr_s := toml_m.A("example").S(1)
      substr_y := []byte(substr_s)
      if ! bytes.Contains(example_y, substr_y) {
         println(index_s)
      }
   }
}
