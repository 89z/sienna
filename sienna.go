package sienna

import (
   "fmt"
   "io"
   "math"
   "path/filepath"
)

func numberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " G"}[n2]
}

type Path struct {
   Base string
   Dir string
   Ext string
   Join string
}

func NewPath(a ...string) Path {
   s := filepath.Join(a...)
   return Path{
      filepath.Base(s),
      filepath.Dir(s),
      filepath.Ext(s),
      s,
   }
}
