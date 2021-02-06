package main

import (
   "bytes"
   "github.com/89z/x"
   "github.com/pelletier/go-toml"
   "io/ioutil"
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

func main() {
   if len(os.Args) != 2 {
      println(`front-matter D:\Git`)
      os.Exit(1)
   }
   root := os.Args[1]
   content := path.Join(root, "autumn", "content")
   dir, e := ioutil.ReadDir(content)
   x.Check(e)
   for _, entry := range dir {
      indexPath := path.Join(
         content, entry.Name(), "_index.md",
      )
      index, e := ioutil.ReadFile(indexPath)
      x.Check(e)
      var front frontMatter
      e = toml.Unmarshal(
         bytes.SplitN(index, tomlSep[:], 3)[1], &front,
      )
      x.Check(e)
      if front.Build.List != "" {
         continue
      }
      examplePath := path.Join(root, front.Filename)
      if ! x.IsFile(examplePath) {
         println(indexPath)
         continue
      }
      example, e := ioutil.ReadFile(examplePath)
      x.Check(e)
      sub := []byte(front.Substr)
      if front.Substr == "" || ! bytes.Contains(example, sub) {
         println(indexPath)
      }
   }
}
