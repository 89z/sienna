package main

import (
   "database/sql"
   "fmt"
   "github.com/89z/sienna/visage"
   "net/http"
   "os"
   "os/exec"
   "path/filepath"
   "strings"
   _ "github.com/89z/sienna/msi"
)

type environ map[string][]string

func (e environ) Add(key, value string) {
   e[key] = append(e[key], value)
}

func (e environ) Encode() []string {
   var enc []string
   for key, val := range e {
      join := strings.Join(val, string(os.PathListSeparator))
      enc = append(enc, key + "=" + join)
   }
   return enc
}

func main() {
   if len(os.Args) == 1 {
      println("visage <command> [args]")
      return
   }
   env := make(environ)
   env.Add("ComSpec", `C:\Windows\System32\cmd.exe`)
   env.Add("PATHEXT", ".exe")
   env.Add("PROCESSOR_ARCHITECTURE", "AMD64")
   env.Add("TMP", `C:\Windows\TEMP`)
   for _, pat := range visage.Patterns {
      matches, err := filepath.Glob(`C:\visage\` + pat)
      if err != nil {
         panic(err)
      }
      if matches == nil {
         panic(pat)
      }
      match := matches[0]
      dir, ext := filepath.Dir(match), filepath.Ext(match)
      key := map[string]string{
         ".EXE": "PATH", ".H": "INCLUDE", ".LIB": "LIB",
      }[strings.ToUpper(ext)]
      env.Add(key, dir)
   }
   cmd := exec.Command(os.Args[1], os.Args[2:]...)
   cmd.Stderr = os.Stderr
   cmd.Stdout = os.Stdout
   cmd.Env = env.Encode()
   err := cmd.Run()
   if err != nil {
      panic(err)
   }
}


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

func download() {
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
