package main

import (
   "log"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate_s := os.Args[1]

   e := System("cargo", "new", "rust-deps")
   if e != nil {
      log.Fatal(e)
   }

   os.Chdir("rust-deps")
   toml_m := Map{
      "dependencies": Map{crate_s: ""},
      "package": Map{"name": "rust-deps", "version": "1.0.0"},
   }

   e = TomlEncode("Cargo.toml", toml_m)
   if e != nil {
      log.Fatal(e)
   }

   e = System("cargo", "generate-lockfile")
   if e != nil {
      log.Fatal(e)
   }

   lock_m, e := TomlDecode("Cargo.lock")
   if e != nil {
      log.Fatal(e)
   }

   prev_s := ""
   dep_n := 0
   name_a := lock_m.A("package")

   for n := range name_a {
      name_s := name_a.M(n).S("name")
      if name_s == "rust-deps" {
         continue
      }
      if name_s == crate_s {
         continue
      }
      if name_s == prev_s {
         continue
      }
      println(name_s)
      prev_s = name_s
      dep_n++
   }

   print("\n", dep_n, " deps\n")
   os.Chdir("..")

   e = os.RemoveAll("rust-deps")
   if e != nil {
      log.Fatal(e)
   }
}
