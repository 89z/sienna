package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "net/http"
   "os"
   "sienna"
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
      url_s := "https://api.godoc.org/search?q=" + repo_s + "/"
      println(url_s)
      get_o, e := http.Get(url_s)
      if e != nil {
         return e
      }
      get_m := sienna.Map{}
      e = json.NewDecoder(get_o.Body).Decode(&get_m)
      if e != nil {
         return e
      }
      result_a := get_m.A("results")
      for n := range result_a {
         path_s := result_a.M(n).S("path")
         if ! strings.HasPrefix(path_s, "golang.org/x/") {
            continue
         }
         path_a := strings.Split(path_s, "/")
         if len(path_a) > 4 {
            continue
         }
         dest := strings.ReplaceAll(path_s, "/", "-")
         _, e = httpCopy("https://pkg.go.dev/" + path_s, dest + ".html")
         if e != nil {
            return e
         }
      }
   }
   return nil
}

func httpCopy(in, out string) (int64, error) {
   println(in)
   source, e := http.Get(in)
   if e != nil {
      return 0, e
   }
   dest, e := os.Create(out)
   if e != nil {
      return 0, e
   }
   return dest.ReadFrom(source.Body)
}
