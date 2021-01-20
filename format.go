package sienna

import (
   "github.com/89z/json"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "os"
)

func TomlGetByte(y []byte) (json.Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func TomlGetFile(s string) (json.Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func TomlPutFile(source json.Map, dest string) error {
   y, e := toml.Marshal(source)
   if e != nil {
      return e
   }
   return ioutil.WriteFile(dest, y, os.ModePerm)
}
