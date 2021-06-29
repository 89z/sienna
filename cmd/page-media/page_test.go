package main
import "testing"

var tests = []struct{in, out string}{
   {
      "https://www.youtube.com/watch?v=LxK5Ocehj10",
      "https://i.ytimg.com/vi/LxK5Ocehj10/maxresdefault.jpg",
   }, {
      "https://www.hrp.org.uk/tower-of-london/history-and-stories/tower-of-london-prison/",
      "https://hrp.imgix.net/https%3A%2F%2Fhistoricroyalpalaces.picturepark.com%2FGo%2FxEpDKCt1%2FV%2F5656%2F36?auto=format&s=66637223c3b7674cbbcfede705be9ae9",
   },
}

func TestOpen(t *testing.T) {
   for _, test := range tests {
      items, err := open(test.in)
      if err != nil {
         t.Fatal(err)
      }
      if items[0] != test.out {
         t.Fatal(items)
      }
   }
}
