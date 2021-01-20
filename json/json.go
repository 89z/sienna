package json
import "encoding/json"

func Load(content string) (Map, error) {
   data := []byte(content)
   v := Map{}
   return v, json.Unmarshal(data, &v)
}

type Map map[string]interface{}

func (m Map) M(key string) Map {
   return m[key].(map[string]interface{})
}
