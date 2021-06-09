package main

import (
   "os"
   "os/exec"
   "path/filepath"
)

func main() {
   bin, err := exec.LookPath("crystal")
   if err != nil {
      panic(err)
   }
   // bin
   bin = filepath.Dir(bin)
   if err := os.Setenv("LIB", os.Getenv("LIB") + ";" + bin)
      err != nil {
      panic(err)
   }
   // src
   src := filepath.Join(filepath.Dir(bin), "src")
   if err := os.Setenv("CRYSTAL_PATH", src)
      err != nil {
      panic(err)
   }
   // Run
   cmd := exec.Command("crystal", os.Args[1:]...)
   cmd.Stderr = os.Stderr
   cmd.Stdout = os.Stdout
   cmd.Stdin = os.Stdin
   if err := cmd.Run()
      err != nil {
      panic(err)
   }
}
