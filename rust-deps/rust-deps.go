package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
   "os/exec"
   "sienna"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func system(command ...string) error {
   name, arg := command[0], command[1:]
   o := exec.Command(name, arg...)
   o.Stderr = os.Stderr
   return o.Run()
}

func tomlDecode(s string) (Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func tomlEncode(s string, m Map) error {
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   defer o.Close()
   return toml.NewEncoder(o).Encode(m)
}

type Map = sienna.Map
