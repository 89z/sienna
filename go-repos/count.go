package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "net/http"
   "sienna/assert"
   "sort"
)

type Repo struct {
   Count int
   Path string
}

var repo_a []Repo

func Count() error {
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
      get_m := assert.Map{}
      json.NewDecoder(get_o.Body).Decode(&get_m)
      result_a := get_m.A("results")
      len_n := len(result_a)
      repo_a = append(repo_a, Repo{len_n, repo_s})
   }
   sort.Slice(repo_a, func(n, n2 int) bool {
      return repo_a[n].Count < repo_a[n2].Count
   })
   for _, o := range repo_a {
      println(o.Count, o.Path)
   }
   return nil
}
