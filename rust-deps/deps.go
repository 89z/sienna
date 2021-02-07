package main

import (
   "github.com/89z/x"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "os"
)

var (
   dep int
   prev string
)

type m map[string]interface{}

type cargoLock struct{
   Package []struct{
      Name string
   }
}

var lock cargoLock

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate := os.Args[1]
   e := x.Command("cargo", "new", "rust-deps").Run()
   x.Check(e)
   e = os.Chdir("rust-deps")
   x.Check(e)
   data, e := toml.Marshal(m{
      "dependencies": m{crate: ""},
      "package": m{"name": "rust-deps", "version": "1.0.0"},
   })
   x.Check(e)
   e = ioutil.WriteFile("Cargo.toml", data, os.ModePerm)
   x.Check(e)
   e = x.Command("cargo", "generate-lockfile").Run()
   x.Check(e)
   data, e = ioutil.ReadFile("Cargo.lock")
   x.Check(e)
   e = toml.Unmarshal(data, &lock)
   x.Check(e)
   for _, pack := range lock.Package {
      if pack.Name == "rust-deps" {
         continue
      }
      if pack.Name == crate {
         continue
      }
      if pack.Name == prev {
         continue
      }
      println(pack.Name)
      prev = pack.Name
      dep++
   }
   print("\n", dep, " deps\n")
   e = os.Chdir("..")
   x.Check(e)
   e = os.RemoveAll("rust-deps")
   x.Check(e)
}
