package main

import (
   "github.com/89z/x"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "os"
   "path"
)

var (
   dep int
   lock cargoLock
   prev string
)

type cargoLock struct{
   Package []struct{
      Name string
   }
}

type m map[string]interface{}

func main() {
   name := "rust-deps"
   if len(os.Args) != 2 {
      println(name, "<crate>")
      os.Exit(1)
   }
   crate := os.Args[1]
   e := x.Command("cargo", "new", name).Run()
   x.Check(e)
   data, e := toml.Marshal(m{
      "dependencies": m{crate: ""},
      "package": m{"name": name, "version": "1.0.0"},
   })
   x.Check(e)
   e = ioutil.WriteFile(path.Join(name, "Cargo.toml"), data, 0)
   x.Check(e)
   cmd := x.Command("cargo", "generate-lockfile")
   cmd.Dir = name
   e = cmd.Run()
   x.Check(e)
   data, e = ioutil.ReadFile(path.Join(name, "Cargo.lock"))
   x.Check(e)
   e = toml.Unmarshal(data, &lock)
   x.Check(e)
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
   x.Check(e)
}
