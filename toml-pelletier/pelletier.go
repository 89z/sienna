package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Slice []interface{}
type Map map[string]interface{}

func main() {
   m := Map{
      "Winter": Slice{
         Map{"December": Slice{Map{"Sunday": 6}, Map{"Monday": 7}}},
         Map{"January": Slice{Map{"Sunday": 6}, Map{"Monday": 7}}},
      },
      "Spring": Slice{
         Map{"March": Slice{Map{"Sunday": 6}, Map{"Monday": 7}}},
         Map{"April": Slice{Map{"Sunday": 6}, Map{"Monday": 7}}},
      },
   }
   e := toml.NewEncoder(os.Stdout).ArraysWithOneElementPerLine(true).Encode(m)
   if e != nil {
      log.Fatal(e)
   }
}
