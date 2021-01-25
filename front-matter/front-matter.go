package main

import (
   "bytes"
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "io/ioutil"
   "os"
   "path"
)

var tomlSep = []byte{'+', '+', '+', '\n'}

func main() {
   if len(os.Args) != 2 {
      println(`front-matter D:\Git`)
      os.Exit(1)
   }
   root := os.Args[1]
   content := path.Join(root, "autumn", "content")
   e := os.Chdir(content)
   x.Check(e)
   dir, e := ioutil.ReadDir(".")
   x.Check(e)
   for _, entry := range dir {
      index_s := path.Join(entry.Name(), "_index.md")
      index_y, e := ioutil.ReadFile(index_s)
      x.Check(e)
      data := bytes.SplitN(index_y, tomlSep, 3)[1]
      front, e := toml.LoadBytes(data)
      x.Check(e)
      if front["_build"] != nil {
         continue
      }
      example := front.A("example")
      exFile := path.Join(root, example.S(0))
      if ! x.IsFile(exFile) {
         println(index_s)
         continue
      }
      example_y, e := ioutil.ReadFile(exFile)
      x.Check(e)
      substr_s := example.S(1)
      substr_y := []byte(substr_s)
      if ! bytes.Contains(example_y, substr_y) {
         println(index_s)
      }
   }
}
