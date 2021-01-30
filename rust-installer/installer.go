package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "github.com/pelletier/go-toml"
   "os"
   "path"
)

const source = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

type distChannel struct{
   Pkg struct{
      Cargo target
      RustStd target `toml:"rust-std"`
      Rustc target
   }
}

type target struct{
   Target struct{
      X8664PcWindowsGnu struct{
         XzUrl string `toml:"xz_url"`
      } `toml:"x86_64-pc-windows-gnu"`
   }
}

var tarXz = archiver.TarXz{
   Tar: &archiver.Tar{OverwriteExisting: true, StripComponents: 2},
}
