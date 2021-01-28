package main
import "testing"

func TestOpen(t *testing.T) {
   out, e := open("https://www.youtube.com/watch?v=LxK5Ocehj10")
   if e != nil {
      t.Error(e)
   }
   if out != "https://i.ytimg.com/vi/LxK5Ocehj10/maxresdefault.jpg" {
      t.Error(out)
   }
}
