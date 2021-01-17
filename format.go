package sienna

import (
   "encoding/json"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "net/http"
   "os"
)

func JsonGetHttp(s string) (Map, error) {
   println(s)
   o, e := http.Get(s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   return m, json.NewDecoder(o.Body).Decode(&m)
}

func TomlGetByte(y []byte) (Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func TomlGetFile(s string) (Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func TomlPutFile(source Map, dest string) error {
   y, e := toml.Marshal(source)
   if e != nil {
      return e
   }
   return ioutil.WriteFile(dest, y, os.ModePerm)
}

type Map map[string]interface{}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}

func (m Map) S(s string) string {
   return m[s].(string)
}

type Slice []interface{}

func (a Slice) M(n int) Map {
   return a[n].(map[string]interface{})
}

func (a Slice) S(n int) string {
   return a[n].(string)
}
