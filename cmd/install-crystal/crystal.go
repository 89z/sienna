package main

import (
   "os"
   "path/filepath"
)

func getCreate(get, create string) error {
   _, err := os.Stat(create)
   if err == nil {
      fmt.Println(invert, "Exist", reset, create)
      return nil
   }
   fmt.Println(invert, "Get", reset, get)
   res, err := http.Get(get)
   if err != nil { return err }
   defer res.Body.Close()
   os.Mkdir(filepath.Dir(create), os.ModeDir)
   file, err := os.Create(create)
   if err != nil { return err }
   defer file.Close()
   _, err = file.ReadFrom(res.Body)
   return err
}

func main() {
   // cache
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna/crystal")
   // get
   err = getCreate(
      "https://codeload.github.com/crystal-lang/crystal/zip/refs/heads/master",
      filepath.Join(cache, "crystal-master.zip"),
   )
   if err != nil {
      panic(err)
   }
   // create
   create := filepath.Join(cache, "crystal.zip")
   // get
   getCreate(
   https://github.com/crystal-lang/crystal/suites/2754559253/artifacts/60912181
   )
}
