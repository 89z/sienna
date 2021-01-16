package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "net/http"
   "sienna"
   "sort"
)

func count() error {
   var repo_a []repo
   for repo_s, repo_o := range repos.ByImportPath {
      if ! repo_o.ShowOnDashboard() {
         continue
      }
      println(repo_s)
      url_s := "https://api.godoc.org/search?q=" + repo_s + "/"
      get_o, e := http.Get(url_s)
      if e != nil {
         return e
      }
      get_m := sienna.Map{}
      json.NewDecoder(get_o.Body).Decode(&get_m)
      result_a := get_m.A("results")
      len_n := len(result_a)
      repo_a = append(repo_a, repo{len_n, repo_s})
   }
   sort.Slice(repo_a, func(n, n2 int) bool {
      return repo_a[n].count < repo_a[n2].count
   })
   for _, o := range repo_a {
      println(o.count, o.path)
   }
   return nil
}

type repo struct {
   count int
   path string
}
