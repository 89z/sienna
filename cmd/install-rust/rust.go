package main

import (
   "net/http"
   "os"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func getCreate(get, create string) error {
   _, err := os.Stat(create)
   if err == nil {
      println(invert, "Exist", reset, create)
      return nil
   }
   println(invert, "Get", reset, get)
   res, err := http.Get(get)
   if err != nil { return err }
   defer res.Body.Close()
   file, err := os.Create(create)
   if err != nil { return err }
   defer file.Close()
   {
      _, err := file.ReadFrom(res.Body)
      return err
   }
}
