package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "net/http"
   "sienna/assert"
   "strings"
)

var justify_m = map[string]string{
   "golang.org/x/arch": "",
   "golang.org/x/benchmarks": "",
   "golang.org/x/blog": "",
   "golang.org/x/build": "",
   "golang.org/x/crypto": "",
   "golang.org/x/exp": "golang.org/x/exp/utf8string",
   "golang.org/x/image": "",
   "golang.org/x/mobile": "",
   "golang.org/x/mod": "",
   "golang.org/x/net": "golang.org/x/net/html",
   "golang.org/x/oauth2": "",
   "golang.org/x/perf": "",
   "golang.org/x/sync": "",
   "golang.org/x/sys": "golang.org/x/sys/windows",
   "golang.org/x/text": "golang.org/x/text/cases",
   "golang.org/x/time": "",
   "golang.org/x/tools": "",
   "golang.org/x/website": "",
}

func BadPath(s string) bool {
   if ! strings.HasPrefix(s, "golang.org/x/") {
      return true
   }
   if strings.Contains(s, "/cmd/") {
      return true
   }
   if strings.Contains(s, "/vendor/") {
      return true
   }
   return false
}

func Download() error {
   count_n := 0
   for repo_s, repo_o := range repos.ByImportPath {
      if ! repo_o.ShowOnDashboard() {
         continue
      }
      if justify_m[repo_s] == "" {
         continue
      }
      url_s := "https://api.godoc.org/search?q=" + repo_s + "/"
      println(url_s)
      get_o, e := http.Get(url_s)
      if e != nil {
         return e
      }
      get_m := assert.Map{}
      json.NewDecoder(get_o.Body).Decode(&get_m)
      result_a := get_m.A("results")
      for n := range result_a {
         path_s := result_a.M(n).S("path")
         if BadPath(path_s) {
            continue
         }
         println("  ", path_s)
         count_n++
      }
   }
   println(count_n)
   return nil
}
