package json

import (
   "encoding/json"
   "net/http"
)

type Map map[string]interface{}

func Load(content string) (Map, error) {
   data := []byte(content)
   m := Map{}
   return m, json.Unmarshal(data, &m)
}

func LoadHttp(url string) (Map, error) {
   resp, err := http.Get(url)
   if err != nil {
      return nil, err
   }
   m := Map{}
   return m, json.NewDecoder(resp.Body).Decode(&m)
}

func (m Map) A(key string) Slice {
   return m[key].([]interface{})
}

func (m Map) M(key string) Map {
   return m[key].(map[string]interface{})
}

func (m Map) N(key string) float64 {
   return m[key].(float64)
}

func (m Map) S(key string) string {
   return m[key].(string)
}

type Slice []interface{}

func (s Slice) M(index int) Map {
   return s[index].(map[string]interface{})
}
