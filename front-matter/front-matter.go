package main

import (
   "bytes"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
   "os"
   "path"
)

var tomlSep = [4]byte{'+', '+', '+', '\n'}

type frontMatter struct{
   Build struct{
      List string
   } `toml:"_build"`
   Filename string
   Substr string
}

func unmarshal(file string, v interface{}) error {
   index, e := ioutil.ReadFile(file)
   if e != nil {
      return e
   }
   return toml.Unmarshal(
      bytes.SplitN(index, tomlSep[:], 3)[1], v,
   )
}

func main() {
   if len(os.Args) != 2 {
      println(`front-matter D:\Git`)
      os.Exit(1)
   }
   root := os.Args[1]
   content := path.Join(root, "autumn", "content")
   dir, e := ioutil.ReadDir(content)
   if e != nil {
      log.Fatal(e)
   }
   for _, entry := range dir {
      indexPath := path.Join(
         content, entry.Name(), "_index.md",
      )
      var front frontMatter
      e = unmarshal(indexPath, &front)
      if e != nil {
         log.Fatal(e)
      }
      if front.Build.List != "" {
         continue
      }
      examplePath := path.Join(root, front.Filename)
      _, err := os.Stat(examplePath)
      if err != nil {
         println(indexPath)
         continue
      }
      example, e := ioutil.ReadFile(examplePath)
      if e != nil {
         log.Fatal(e)
      }
      sub := []byte(front.Substr)
      if front.Substr == "" || ! bytes.Contains(example, sub) {
         println(indexPath)
      }
   }
}
