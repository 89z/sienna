package main

import (
   "github.com/mholt/archiver/v3"
   "log"
)

const channel = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func unarchive(file, dir string) error {
   tar := &archiver.Tar{OverwriteExisting: true, StripComponents: 2}
   println("EXTRACT", file)
   xz := archiver.TarXz{Tar: tar}
   return xz.Unarchive(file, dir)
}
