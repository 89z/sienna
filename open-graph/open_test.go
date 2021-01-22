package main
import "testing"

type test struct {
   in string
   out string
}

var tests = []test{
   {
      "https://www.youtube.com/watch?v=LxK5Ocehj10",
      "https://i.ytimg.com/vi/LxK5Ocehj10/maxresdefault.jpg",
   },
}

func TestOpen(t *testing.T) {
   for _, each := range tests {
      out, e := open(each.in)
      if e != nil {
         t.Error(e)
      }
      if out != each.out {
         t.Error(out)
      }
   }
}
