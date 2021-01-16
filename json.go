package sienna

import (
   "fmt"
   "io"
   "math"
   "net/http"
   "os"
)

type oMap map[string]interface{}

func (m oMap) a(s string) slice {
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
