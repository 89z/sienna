package main

import (
   "github.com/89z/sienna"
   "golang.org/x/build/repos"
   "os"
   "strings"
)

var bad_repo = map[string]bool{
   "golang.org/x/build": true,
   "golang.org/x/crypto": true,
   "golang.org/x/oauth2": true,
   "golang.org/x/tools": true,
}

func download() error {
   os.Mkdir("x", os.ModeDir)
   os.Chdir("x")
   for repo_s, repo_o := range repos.ByImportPath {
      if ! repo_o.ShowOnDashboard() {
         continue
      }
      if bad_repo[repo_s] {
         continue
      }
      url := "https://api.godoc.org/search?q=" + repo_s + "/"
      get, e := sienna.JsonGetHttp(url)
      if e != nil {
         return e
      }
      result_a := get.A("results")
      for n := range result_a {
         path := result_a.M(n).S("path")
         if ! strings.HasPrefix(path, "golang.org/x/") {
            continue
         }
         path_a := strings.Split(path, "/")
         if len(path_a) > 4 {
            continue
         }
         dest := strings.ReplaceAll(path, "/", "-")
         _, e = sienna.HttpCopy("https://pkg.go.dev/" + path, dest + ".html")
         if e != nil {
            return e
         }
      }
   }
   return nil
}
