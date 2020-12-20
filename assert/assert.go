package assert
type Map map[string]interface{}
type Slice []interface{}

func (a Slice) S(n int) string {
   return a[n].(string)
}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}

func (m Map) S(s string) string {
   return m[s].(string)
}
