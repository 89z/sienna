package main

import (
   "github.com/89z/sienna/msi"
   "github.com/89z/sienna/visage"
   "os"
   "path/filepath"
)

func install() {
   // cache
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "visage")
   // create
   for _, pack := range visage.Packages {
      for _, payload := range pack.Payloads {
         create := filepath.Join(cache, pack.ID, payload.Filename)
         if filepath.Ext(payload.Filename) == ".vsix" {
            var extract visage.Archive
            println(invert, "Zip", reset, payload.Filename)
            err := extract.Zip(create, `C:\visage`)
            if err != nil {
               panic(err)
            }
            continue
         }
         println(invert, "MSI", reset, payload.Filename)
         err := msi.InstallProduct(create, `Action=Admin TargetDir=C:\visage`)
         if err != nil {
            panic(err)
         }
      }
   }
}
