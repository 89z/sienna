package release

type Map map[string]interface{}
type Slice []interface{}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}

func (m Map) S(key_s string) (string, bool) {
   val_s, b := m[key_s]
   return val_s.(string), b
}
