package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "net/http"
   "sienna/assert"
)

func Count() error {
   for repo_s, repo_o := range repos.ByImportPath {
      if ! repo_o.ShowOnDashboard() {
         continue
      }
      url_s := "https://api.godoc.org/search?q=" + repo_s + "/"
      get_o, e := http.Get(url_s)
      if e != nil {
         return e
      }
      get_m := assert.Map{}
      json.NewDecoder(get_o.Body).Decode(&get_m)
      result_a := get_m.A("results")
      len_n := len(result_a)
      println(len_n, repo_s)
   }
   return nil
}
