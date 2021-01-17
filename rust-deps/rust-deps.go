package main

import (
   "log"
   "os"
   "sienna"
)

var (
   dep int
   prev string
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

type Map = sienna.Map

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate := os.Args[1]
   e := sienna.System("cargo", "new", "rust-deps")
   check(e)
   os.Chdir("rust-deps")
   e = sienna.TomlPutFile(
      Map{
         "dependencies": Map{crate: ""},
         "package": Map{"name": "rust-deps", "version": "1.0.0"},
      },
      "Cargo.toml",
   )
   check(e)
   e = sienna.System("cargo", "generate-lockfile")
   check(e)
   lock, e := sienna.TomlGetFile("Cargo.lock")
   check(e)
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
   check(e)
}
