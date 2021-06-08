package main

import (
   "archive/zip"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
   "strings"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func netrc(addr string) (*http.Request, error) {
   home, err := os.UserHomeDir()
   if err != nil { return nil, err }
   file, err := os.Open(home + "/_netrc")
   if err != nil { return nil, err }
   defer file.Close()
   var login, pass string
   if _, err := fmt.Fscanf(
      file, "default login %v password %v", &login, &pass,
   ); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil { return nil, err }
   req.SetBasicAuth(login, pass)
   return req, nil
}

func httpGet(req *http.Request, create string) error {
   _, err := os.Stat(create)
   if err == nil {
      fmt.Println(invert, "Exist", reset, create)
      return nil
   }
   fmt.Println(invert, "Get", reset, req.URL)
   res, err := new(http.Client).Do(req)
   if err != nil { return err }
   defer res.Body.Close()
   if err := os.MkdirAll(filepath.Dir(create), os.ModeDir); err != nil {
      return err
   }
   file, err := os.Create(create)
   if err != nil { return err }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

type archive struct { strip int }

func (a archive) extractZip(source, dest string) error {
   read, err := zip.OpenReader(source)
   if err != nil { return err }
   defer read.Close()
   for _, file := range read.File {
      if file.Mode().IsDir() { continue }
      name := a.stripPath(dest, file.Name)
      if name == "" { continue }
      if err := os.MkdirAll(filepath.Dir(name), os.ModeDir); err != nil {
         return err
      }
      open, err := file.Open()
      if err != nil { return err }
      create, err := os.Create(name)
      if err != nil { return err }
      defer create.Close()
      if _, err := create.ReadFrom(open); err != nil {
         return err
      }
   }
   return nil
}

func (a archive) stripPath(left, right string) string {
   split := strings.SplitN(right, "/", a.strip + 1)
   if len(split) <= a.strip { return "" }
   return filepath.Join(left, split[a.strip])
}


func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "crystal")
   create := filepath.Join(cache, "crystal-master.zip")
   req, err := http.NewRequest(
      "GET",
      "https://codeload.github.com/crystal-lang/crystal/zip/refs/heads/master",
      nil,
   )
   if err != nil {
      panic(err)
   }
   if err := httpGet(req, create); err != nil {
      panic(err)
   }
   arc := archive{2}
   fmt.Println(invert, "Extract", reset, create)
   if err := arc.extractZip(create, `C:\crystal`); err != nil {
      panic(err)
   }
   create = filepath.Join(cache, "crystal.zip")
   if req, err := netrc(
      "https://api.github.com" +
      "/repos/crystal-lang/crystal/actions/artifacts/61069508/zip",
   ); err != nil {
      panic(err)
   } else if err := httpGet(req, create); err != nil {
      panic(err)
   }
   arc.strip = 0
   fmt.Println(invert, "Extract", reset, create)
   if err := arc.extractZip(create, `C:\crystal\bin`); err != nil {
      panic(err)
   }
}
