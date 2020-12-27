package main

import (
   "fmt"
   "golang.org/x/build/repos"
)

func main() {
   for s, o := range repos.ByImportPath {
      if o.CoordinatorCanBuild {
         fmt.Println(s)
      }
   }
}
