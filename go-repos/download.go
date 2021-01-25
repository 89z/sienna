package main

import (
   "github.com/89z/x"
   "github.com/89z/x/json"
   "golang.org/x/build/repos"
   "strings"
)

var badRepo = map[string]bool{
   "golang.org/x/build": true,
   "golang.org/x/crypto": true,
   "golang.org/x/oauth2": true,
   "golang.org/x/tools": true,
}

func download() error {
   for importPath, info := range repos.ByImportPath {
      if ! info.ShowOnDashboard() {
         continue
      }
      if badRepo[importPath] {
         continue
      }
      get, e := json.LoadHttp(
         "https://api.godoc.org/search?q=" + importPath + "/",
      )
      if e != nil {
         return e
      }
      results := get.A("results")
      for n := range results {
         path := results.M(n).S("path")
         if ! strings.HasPrefix(path, "golang.org/x/") {
            continue
         }
         if strings.Count(path, "/") > 3 {
            continue
         }
         _, e = x.HttpCopy("https://pkg.go.dev/" + path, path)
         if e != nil {
            return e
         }
      }
   }
   return nil
}
