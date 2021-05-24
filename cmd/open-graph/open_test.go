package main
import "testing"

func TestOpen(t *testing.T) {
   out, err := open("https://www.youtube.com/watch?v=LxK5Ocehj10")
   if err != nil {
      t.Error(err)
   }
   if out != "https://i.ytimg.com/vi/LxK5Ocehj10/maxresdefault.jpg" {
      t.Error(out)
   }
}
