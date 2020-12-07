package main

import (
   "encoding/json"
   "fmt"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("toml <json file>")
      os.Exit(1)
   }

   path_s := os.Args[1]
   open_o, e := os.Open(path_s)
   if e != nil {
      log.Fatal(e)
   }

   m := map[string]interface{}{}
   json.NewDecoder(open_o).Decode(&m)
   fmt.Println(m)
   create_o, e := os.Create("z.toml")
   if e != nil {
      log.Fatal(e)
   }

   e = toml.NewEncoder(create_o).QuoteMapKeys(true).Encode(m)
   if e != nil {
      log.Fatal(e)
   }
}
