package sienna

import (
   "fmt"
   "io"
   "math"
   "net/http"
   "os"
   "os/exec"
   "path/filepath"
)

func HttpCopy(in, out string) (int64, error) {
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

func IsFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

func numberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " G"}[n2]
}

func System(command ...string) error {
   name, arg := command[0], command[1:]
   o := exec.Command(name, arg...)
   o.Stderr = os.Stderr
   o.Stdout = os.Stdout
   println(command)
   return o.Run()
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
