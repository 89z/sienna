package sienna

import (
   "fmt"
   "io"
   "math"
   "os"
   "os/exec"
   "path/filepath"
)

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
