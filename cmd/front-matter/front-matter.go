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
   index, err := os.ReadFile(file)
   if err != nil { return err }
   return toml.Unmarshal(
      bytes.SplitN(index, tomlSep[:], 3)[1], v,
   )
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
      indexPath := filepath.Join(
         content, each.Name(), "_index.md",
      )
      var front frontMatter
      err = unmarshal(indexPath, &front)
      if err != nil {
         panic(err)
      }
      if front.Build.List != "" { continue }
      examplePath := filepath.Join(root, front.Filename)
      _, err := os.Stat(examplePath)
      if err != nil {
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
