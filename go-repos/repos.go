package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "log"
   "net/http"
   "sienna/assert"
)

func BadRepo(s string) bool {
   if s == "golang.org/x/arch" {
      return true
   }
   if s == "golang.org/x/benchmarks" {
      return true
   }
   if s == "golang.org/x/blog" {
      return true
   }
   if s == "golang.org/x/mod" {
      return true
   }
   if s == "golang.org/x/oauth2" {
      return true
   }
   if s == "golang.org/x/perf" {
      return true
   }
   if s == "golang.org/x/sync" {
      return true
   }
   if s == "golang.org/x/website" {
      return true
   }
   return false
}

func main() {
   for repo_s, repo_o := range repos.ByImportPath {
      if repo_o.ShowOnDashboard() {
         if BadRepo(repo_s) {
            continue
         }
         url_s := "https://api.godoc.org/search?q=" + repo_s + "/"
         println(url_s)
         get_o, e := http.Get(url_s)
         if e != nil {
            log.Fatal(e)
         }
         get_m := assert.Map{}
         json.NewDecoder(get_o.Body).Decode(&get_m)
         result_a := get_m.A("results")
         if len(result_a) == 100 {
            continue
         }
         for n := range result_a {
            path_s := result_a.M(n).S("path")
            println("  ", path_s)
         }
      }
   }
}
