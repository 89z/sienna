package net

import (
   "fmt"
   "io"
   "math"
   "net/http"
   "os"
)

type Progress struct {
   Parent io.Reader
   Total float64
}

func (o *Progress) Read(y []byte) (int, error) {
   n, e := o.Parent.Read(y)
   if e != nil {
      fmt.Println()
   } else {
      o.Total += float64(n)
      fmt.Printf("READ %9s\r", NumberFormat(o.Total))
   }
   return n, e
}

func Copy(url_s, path_s string) error {
   get_o, e := http.Get(url_s)
   if e != nil {
      return e
   }
   create_o, e := os.Create(path_s)
   if e != nil {
      return e
   }
   fmt.Println("GET", url_s)
   prog_o := &Progress{get_o.Body, 0}
   n, e := io.Copy(create_o, prog_o)
   if e != nil {
      return fmt.Errorf("%v %v", n, e)
   }
   return nil
}

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " G"}[n2]
}
