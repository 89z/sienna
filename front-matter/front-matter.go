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

func tomlDecode(y []byte) (sienna.Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   if len(os.Args) != 2 {
      println(`front-matter D:\Git`)
      os.Exit(1)
   }
   root := os.Args[1]
   content := path.Join(root, "autumn", "content")
   e := os.Chdir(content)
   check(e)
   dir, e := ioutil.ReadDir(".")
   check(e)
   for _, entry := range dir {
      index_s := path.Join(entry.Name(), "_index.md")
      index_y, e := ioutil.ReadFile(index_s)
      check(e)
      toml_y := bytes.SplitN(index_y, toml_sep, 3)[1]
      toml_m, e := tomlDecode(toml_y)
      check(e)
      if toml_m["_build"] != nil {
         continue
      }
      example := toml_m.A("example")
      exFile := path.Join(root, example.S(0))
      if ! isFile(exFile) {
         println(index_s)
         continue
      }
      example_y, e := ioutil.ReadFile(exFile)
      check(e)
      substr_s := example.S(1)
      substr_y := []byte(substr_s)
      if ! bytes.Contains(example_y, substr_y) {
         println(index_s)
      }
   }
}
