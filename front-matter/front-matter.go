package main
import "log"

var toml_sep = []byte{'+', '+', '+', '\n'}

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}
