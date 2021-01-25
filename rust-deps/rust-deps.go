package main

import (
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "os"
)

var (
   dep int
   prev string
)

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate := os.Args[1]
   e := x.System("cargo", "new", "rust-deps")
   x.Check(e)
   os.Chdir("rust-deps")
   e = toml.DumpFile(
      "Cargo.toml",
      x.Map{
         "dependencies": x.Map{crate: ""},
         "package": x.Map{"name": "rust-deps", "version": "1.0.0"},
      },
   )
   x.Check(e)
   e = x.System("cargo", "generate-lockfile")
   x.Check(e)
   lock, e := toml.LoadFile("Cargo.lock")
   x.Check(e)
   packages := lock.A("package")
   for n := range packages {
      name := packages.M(n).S("name")
      if name == "rust-deps" {
         continue
      }
      if name == crate {
         continue
      }
      if name == prev {
         continue
      }
      println(name)
      prev = name
      dep++
   }
   print("\n", dep, " deps\n")
   os.Chdir("..")
   e = os.RemoveAll("rust-deps")
   x.Check(e)
}
