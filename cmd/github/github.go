package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "os"
)

func netrc() (string, string, error) {
   home, e := os.UserHomeDir()
   if e != nil { return "", "", e }
   f, e := os.Open(home + "/_netrc")
   if e != nil { return "", "", e }
   defer f.Close()
   var login, pass string
   fmt.Fscanf(f, "default login %v password %v", &login, &pass)
   return login, pass, nil
}

type repoSearch struct {
   *http.Request
   lang map[string]bool
}

func newRepoSearch(stars, date string) (repoSearch, error) {
   login, pass, e := netrc()
   if e != nil {
      return repoSearch{}, e
   }
   req, e := http.NewRequest(
      "GET", "https://api.github.com/search/repositories", nil,
   )
   if e != nil {
      return repoSearch{}, e
   }
   val := req.URL.Query()
   val.Set("per_page", "100")
   val.Set("q", fmt.Sprintf("stars:>=%v pushed:>=%v", stars, date))
   req.URL.RawQuery = val.Encode()
   req.SetBasicAuth(login, pass)
   return repoSearch{
      req, make(map[string]bool),
   }, nil
}

func (req repoSearch) page(num int) error {
   val := req.URL.Query()
   val.Set("page", fmt.Sprint(num))
   req.URL.RawQuery = val.Encode()
   res, e := http.DefaultClient.Do(req.Request)
   if e != nil { return e }
   defer res.Body.Close()
   var result struct {
      Items []struct {
         HTML_URL string
         Language string
         Stargazers_Count int
      }
   }
   json.NewDecoder(res.Body).Decode(&result)
   for _, item := range result.Items {
      if item.Language == "" { continue }
      if req.lang[item.Language] { continue }
      fmt.Printf(
         "%-10v %v %v\n", item.Language, item.Stargazers_Count, item.HTML_URL,
      )
      req.lang[item.Language] = true
   }
   return nil
}
