package main

import (
   "bytes"
   "github.com/pelletier/go-toml"
   "os"
   "path/filepath"
)

var tomlSep = [4]byte{'+', '+', '+', '\n'}

type frontMatter struct {
   Build struct {
      List string
   } `toml:"_build"`
   Filename string
   Substr string
}

func unmarshal(file string, v interface{}) error {
   index, e := os.ReadFile(file)
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
   content := filepath.Join(root, "autumn", "content")
   dir, e := os.ReadDir(content)
   if e != nil {
      panic(e)
   }
   for _, each := range dir {
      indexPath := filepath.Join(
         content, each.Name(), "_index.md",
      )
      var front frontMatter
      e = unmarshal(indexPath, &front)
      if e != nil {
         panic(e)
      }
      if front.Build.List != "" {
         continue
      }
      examplePath := filepath.Join(root, front.Filename)
      _, err := os.Stat(examplePath)
      if err != nil {
         println(indexPath)
         continue
      }
      example, e := os.ReadFile(examplePath)
      if e != nil {
         panic(e)
      }
      sub := []byte(front.Substr)
      if front.Substr == "" || ! bytes.Contains(example, sub) {
         println(indexPath)
      }
   }
}
