package main

import (
   "encoding/json"
   "github.com/89z/x"
   "golang.org/x/build/repos"
   "net/http"
   "os"
   "strings"
)

var badRepo = map[string]bool{
   "golang.org/x/build": true,
   "golang.org/x/oauth2": true,
   "golang.org/x/tools": true,
}

type search struct {
   Results []struct {
      Path string
   }
}

func download() error {
   var godoc search
   for importPath, info := range repos.ByImportPath {
      if ! info.ShowOnDashboard() {
         continue
      }
      if badRepo[importPath] {
         continue
      }
      get, e := http.Get("https://api.godoc.org/search?q=" + importPath + "/")
      if e != nil {
         return e
      }
      e = json.NewDecoder(get.Body).Decode(&godoc)
      if e != nil {
         return e
      }
      for _, result := range godoc.Results {
         if ! strings.HasPrefix(result.Path, "golang.org/x/") {
            continue
         }
         if strings.Count(result.Path, "/") > 3 {
            continue
         }
         _, e = x.Copy("https://pkg.go.dev/" + result.Path, result.Path)
         if os.IsExist(e) {
            println(x.ColorCyan("Exist"), result.Path)
         } else if e != nil {
            return e
         }
      }
   }
   return nil
}
