package main

import (
   "os"
   "path/filepath"
)

func main() {
   // cache
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna/crystal")
   // create
   create := filepath.Join(cache, "crystal-master.zip")
   // req
   // https://codeload.github.com/crystal-lang/crystal/zip/refs/heads/master
   // create
   create = filepath.Join(cache, "crystal.zip")
   // get
   // github.com/crystal-lang/crystal/suites/2754559253/artifacts/60912181
   println(create)
}
