package main

import (
   "github.com/pelletier/go-toml"
   "os"
   "os/exec"
)

func system(command ...string) error {
   name, arg := command[0], command[1:]
   o := exec.Command(name, arg...)
   o.Stderr = os.Stderr
   return o.Run()
}

func tomlDecode(s string) (oMap, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func tomlEncode(s string, m oMap) error {
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   defer o.Close()
   return toml.NewEncoder(o).Encode(m)
}

type oMap map[string]interface{}

func (m oMap) a(s string) slice {
   return m[s].([]interface{})
}

func (m oMap) s(key string) string {
   return m[key].(string)
}

type slice []interface{}

func (a slice) m(n int) oMap {
   return a[n].(map[string]interface{})
}
