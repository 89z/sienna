package main

import (
   "bytes"
   "github.com/89z/x"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "os"
   "path"
)

func main() {
   if len(os.Args) != 2 {
      println(`front-matter D:\Git`)
      os.Exit(1)
   }
   root := os.Args[1]
   e := os.Chdir(
      path.Join(root, "autumn", "content"),
   )
   x.Check(e)
   dir, e := ioutil.ReadDir(".")
   x.Check(e)
   for _, entry := range dir {
      indexPath := path.Join(
         entry.Name(), "_index.md",
      )
      index, e := ioutil.ReadFile(indexPath)
      x.Check(e)
      front := new(frontMatter)
      e = toml.Unmarshal(
         bytes.SplitN(index, tomlSep, 3)[1], front,
      )
      x.Check(e)
      if front.Build.List != "" {
         continue
      }
      examplePath := path.Join(
         root, front.Example[0],
      )
      if ! x.IsFile(examplePath) {
         println(indexPath)
         continue
      }
      example, e := ioutil.ReadFile(examplePath)
      x.Check(e)
      substr := []byte(
         front.Example[1],
      )
      if ! bytes.Contains(example, substr) {
         println(indexPath)
      }
   }
}
