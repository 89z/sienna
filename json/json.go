package json
import "encoding/json"

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

func (m Map) M(key string) Map {
   return m[key].(map[string]interface{})
}
