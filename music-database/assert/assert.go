package assert

type Map map[string]interface{}
type Slice []interface{}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}

func (m Map) N(s string) float64 {
   return m[s].(float64)
}

func (m Map) S(s string) string {
   return m[s].(string)
}

func (a Slice) M(n int) Map {
   return a[n].(map[string]interface{})
}
