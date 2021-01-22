package main

import (
   "encoding/json"
   "github.com/89z/x"
   "io/ioutil"
   "log"
   "os"
   "path"
   "regexp"
   "strconv"
   "time"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func findSubmatch(re string, input []byte) string {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return string(a[1])
}
