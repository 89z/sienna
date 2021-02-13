package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
   "os/exec"
   "path"
)

const name = "rust-deps"

type m map[string]interface{}

type lockFile struct {
   Package []struct {
      Name string
   }
}

func cargoLock(v interface{}) error {
   cmd := exec.Command("cargo", "generate-lockfile")
   cmd.Dir = name
   cmd.Stderr = os.Stderr
   cmd.Stdout = os.Stdout
   e := cmd.Run()
   if e != nil {
      return e
   }
   open, e := os.Open(
      path.Join(name, "Cargo.lock"),
   )
   if e != nil {
      return e
   }
   defer open.Close()
   return toml.NewDecoder(open).Decode(v)
}

func cargoToml(crate string) error {
   e := exec.Command("cargo", "new", name).Run()
   if e != nil {
      return e
   }
   create, e := os.Create(
      path.Join(name, "Cargo.toml"),
   )
   if e != nil {
      return e
   }
   defer create.Close()
   return toml.NewEncoder(create).Encode(m{
      "dependencies": m{crate: ""},
      "package": m{"name": name, "version": "1.0.0"},
   })
}

func main() {
   if len(os.Args) != 2 {
      println(name, "<crate>")
      os.Exit(1)
   }
   crate := os.Args[1]
   e := cargoToml(crate)
   if e != nil {
      log.Fatal(e)
   }
   var lock lockFile
   e = cargoLock(&lock)
   if e != nil {
      log.Fatal(e)
   }
   var (
      dep int
      prev string
   )
   for _, pack := range lock.Package {
      if pack.Name == name {
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
   e = os.RemoveAll(name)
   if e != nil {
      log.Fatal(e)
   }
}
