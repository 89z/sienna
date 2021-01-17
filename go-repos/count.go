package main

import (
   "golang.org/x/build/repos"
   "sort"
   "sienna"
)

func count() error {
   var repo_a []repo
   for repo_s, repo_o := range repos.ByImportPath {
      if ! repo_o.ShowOnDashboard() {
         continue
      }
      url := "https://api.godoc.org/search?q=" + repo_s + "/"
      m, e := sienna.JsonGetHttp(url)
      if e != nil {
         return e
      }
      results := m.A("results")
      size := len(results)
      repo_a = append(repo_a, repo{size, repo_s})
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
