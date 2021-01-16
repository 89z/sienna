package sienna

import (
   "fmt"
   "io"
   "math"
   "net/http"
   "os"
)

func httpCopy(in, out string) (int64, error) {
   get, e := http.Get(in)
   if e != nil {
      return 0, e
   }
   dest, e := os.Create(out)
   if e != nil {
      return 0, e
   }
   fmt.Println("GET", in)
   source := progress{get.Body, 0}
   return io.Copy(dest, &source)
}

func numberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " G"}[n2]
}

type progress struct {
   parent io.Reader
   total float64
}

func (o *progress) Read(y []byte) (int, error) {
   n, e := o.parent.Read(y)
   if e != nil {
      fmt.Println()
   } else {
      o.total += float64(n)
      fmt.Printf("READ %9s\r", numberFormat(o.total))
   }
   return n, e
}
