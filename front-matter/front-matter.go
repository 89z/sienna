package main

import (
   "bytes"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
   "os"
)

type Map map[string]interface{}

func (m Map) S(s string) string {
   return m[s].(string)
}

func IsFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

var toml_sep = []byte{'+', '+', '+', '\n'}

func Decode(s string) (string, error) {
   file_y, e := ioutil.ReadFile(s + `\_index.md`)
   if e != nil {
      return "", e
   }
   toml_y := bytes.SplitN(file_y, toml_sep, 3)[1]
   m := Map{}
   e = toml.Unmarshal(toml_y, &m)
   if e != nil {
      return "", e
   }
   if m["example"] == nil {
      return "", nil
   }
   return `D:\Git\` + m.S("example"), nil
}

func main() {
   os.Chdir(`D:\Git\autumn\content`)
   a, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for n := range a {
      name := a[n].Name()
      example, e := Decode(name)
      if e != nil {
         log.Fatal(e)
      }
      if example == "" {
         continue
      }
      if ! IsFile(example) {
         println(name)
      }
   }
}
