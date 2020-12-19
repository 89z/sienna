package main

import (
   "github.com/pelletier/go-toml"
   "os"
   "os/exec"
)

type Map map[string]interface{}
type Slice []interface{}

func (a Slice) M(n int) Map {
   return a[n].(map[string]interface{})
}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}

func (m Map) S(s string) string {
   return m[s].(string)
}

func System(command ...string) error {
   name, arg := command[0], command[1:]
   o := exec.Command(name, arg...)
   o.Stdout = os.Stdout
   return o.Run()
}

func TomlDecode(filename string) (Map, error) {
   o, e := os.Open(filename)
   if e != nil {
      return nil, e
   }
   m := Map{}
   return m, toml.NewDecoder(o).Decode(&m)
}

func TomlEncode(filename string, data Map) error {
   o, e := os.Create(filename)
   if e != nil {
      return e
   }
   return toml.NewEncoder(o).Encode(data)
}
