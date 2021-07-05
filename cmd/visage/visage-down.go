package main

import (
   "database/sql"
   "fmt"
   "github.com/89z/visage"
   "net/http"
   "os"
   "path/filepath"
   _ "github.com/89z/visage/msi"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func getCreate(get, create string) error {
   _, err := os.Stat(create)
   if err == nil {
      fmt.Println(invert, "Exist", reset, create)
      return nil
   }
   res, err := http.Get(get)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := os.MkdirAll(filepath.Dir(create), os.ModeDir); err != nil {
      return err
   }
   file, err := os.Create(create)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func main() {
   // cache
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "visage")
   // create
   create := filepath.Join(cache, "winget-pkgs-master.tar.gz")
   // save
   if err := getCreate(visage.WingetURL, create); err != nil {
      panic(err)
   }
   // get
   winget, err := visage.NewWinGet(create)
   if err != nil {
      panic(err)
   }
   get, err := winget.ChannelURI()
   if err != nil {
      panic(err)
   }
   // create
   create = filepath.Join(cache, "VisualStudio.chman")
   // save
   if err := getCreate(get, create); err != nil {
      panic(err)
   }
   // get
   chman, err := visage.NewChannelMan(create)
   if err != nil {
      panic(err)
   }
   get = chman.VisualStudioURL()
   // create
   create = filepath.Join(cache, "VisualStudio.vsman")
   // save
   if err := getCreate(get, create); err != nil {
      panic(err)
   }
   // get
   vsman, err := visage.NewVisualStudioMan(create)
   if err != nil {
      panic(err)
   }
   for _, pack := range visage.Packages {
      for _, payload := range pack.Payloads {
         get, err := vsman.PayloadURL(pack.ID, payload.Filename)
         if err != nil {
            panic(err)
         }
         create := filepath.Join(cache, pack.ID, payload.Filename)
         if err := getCreate(get, create); err != nil {
            panic(err)
         }
         if filepath.Ext(payload.Filename) == ".vsix" {
            continue
         }
         db, err := sql.Open("msi", create)
         if err != nil {
            panic(err)
         }
         defer db.Close()
         rows, err := db.Query("SELECT Cabinet FROM Media")
         if err != nil {
            panic(err)
         }
         defer rows.Close()
         for rows.Next() {
            var cab string
            err := rows.Scan(&cab)
            if err != nil {
               panic(err)
            }
            if cab == "" {
               continue
            }
            get, err := vsman.PayloadURL(pack.ID, cab)
            if err != nil {
               panic(err)
            }
            create := filepath.Join(cache, pack.ID, cab)
            if err := getCreate(get, create); err != nil {
               panic(err)
            }
         }
      }
   }
}
