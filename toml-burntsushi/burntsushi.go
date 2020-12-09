package main

import (
   "github.com/BurntSushi/toml"
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
   toml.NewEncoder(os.Stdout).Encode(m)
}
