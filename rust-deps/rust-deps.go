package main
import "log"

var (
   dep int
   prev string
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

type Map map[string]interface{}
