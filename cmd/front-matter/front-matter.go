package main

import (
   "bytes"
   "github.com/pelletier/go-toml/v2"
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

func main() {
   if len(os.Args) != 2 {
      println(`front-matter D:\Git`)
      return
   }
   root := os.Args[1]
   content := filepath.Join(root, "autumn", "content")
   dir, err := os.ReadDir(content)
   if err != nil {
      panic(err)
   }
   for _, each := range dir {
      indexPath := filepath.Join(content, each.Name(), "_index.md")
      index, err := os.ReadFile(indexPath)
      if err != nil {
         panic(err)
      }
      index = bytes.SplitN(index, tomlSep[:], 3)[1]
      var front frontMatter
      if err := toml.Unmarshal(index, &front); err != nil {
         panic(err)
      }
      if front.Build.List != "" {
         continue
      }
      if front.Filename == "" {
         println(indexPath)
         continue
      }
      examplePath := filepath.Join(root, front.Filename)
      if _, err := os.Stat(examplePath); err != nil {
         println(indexPath)
         continue
      }
      example, err := os.ReadFile(examplePath)
      if err != nil {
         panic(err)
      }
      sub := []byte(front.Substr)
      if front.Substr == "" || ! bytes.Contains(example, sub) {
         println(indexPath)
      }
   }
}
