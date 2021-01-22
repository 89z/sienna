package main
import "testing"

var tests = []struct{
   in string
   out []string
}{
   "https://www.youtube.com/watch?v=LxK5Ocehj10",
   {"https://i.ytimg.com/vi/LxK5Ocehj10/maxresdefault.jpg"},
}

func TestOpen(t *testing.T) {
   for _, test := range tests {
   }
}
