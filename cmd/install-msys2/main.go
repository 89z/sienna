package main

import (
   "fmt"
   "github.com/89z/sienna"
   "net/http"
   "os"
   "path/filepath"
)

func main() {
   if len(os.Args) != 3 {
      fmt.Println(`install-msys2 query git
install-msys2 sync git.txt`)
      return
   }
   data := newDatabase()
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "msys2")
   for _, db := range []string{
      "/mingw/ucrt64/ucrt64.db",
      "/mingw/x86_64/mingw64.db",
      "/msys/x86_64/msys.db",
   } {
      create := filepath.Join(cache, db)
      _, err := os.Stat(create)
      if err != nil {
         r, err := http.Get(mirror + db)
         if err != nil {
            panic(err)
         }
         defer r.Body.Close()
         f, err := os.Create(create)
         if err != nil {
            panic(err)
         }
         defer f.Close()
         f.ReadFrom(r.Body)
      } else {
         fmt.Println(invert, "Exist", reset, db)
      }
      fs, err := sienna.TarGzMemory(create)
      if err != nil {
         panic(err)
      }
      for _, file := range fs {
         data.scan(file.Data)
      }
   }
   target := os.Args[2]
   if os.Args[1] == "sync" {
      data.sync(target)
      return
   }
   data.query(target)
}
