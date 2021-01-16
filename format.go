package sienna

import (
   "encoding/json"
   "net/http"
)

func JsonFromHttp(s string) (Map, error) {
   println(s)
   o, e := http.Get(s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   return m, json.NewDecoder(o.Body).Decode(&m)
}

func TomlDecode(y []byte) (sienna.Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func tomlDecode(s string) (Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func tomlEncode(s string, m Map) error {
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   defer o.Close()
   return toml.NewEncoder(o).Encode(m)
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
