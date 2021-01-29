package main

import (
   "encoding/json"
   "golang.org/x/build/repos"
   "net/http"
   "sort"
)

func count() error {
   var packs []pack
   for path, value := range repos.ByImportPath {
      if ! value.ShowOnDashboard() {
         continue
      }
      url := "https://api.godoc.org/search?q=" + path + "/"
      println(url)
      get, e := http.Get(url)
      if e != nil {
         return e
      }
      body := new(search)
      e = json.NewDecoder(get.Body).Decode(&body)
      if e != nil {
         return e
      }
      packs = append(packs, pack{
         len(body.Results), path,
      })
   }
   sort.Slice(packs, func(i, j int) bool {
      return packs[i].count < packs[j].count
   })
   for _, each := range packs {
      println(each.count, each.path)
   }
   return nil
}

type pack struct {
   count int
   path string
}
