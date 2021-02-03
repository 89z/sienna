package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "path/filepath"
)

const source = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

var tarXz = archiver.TarXz{
   Tar: &archiver.Tar{OverwriteExisting: true, StripComponents: 2},
}

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

func main() {
   install, e := x.NewInstall("rust")
   x.Check(e)
   cache := filepath.Join(
      install.Cache, filepath.Base(source),
   )
   if ! x.IsFile(cache) {
      _, e = x.Copy(source, cache)
      x.Check(e)
   }
   data, e := ioutil.ReadFile(cache)
   x.Check(e)
   channel := new(distChannel)
   e = toml.Unmarshal(data, channel)
   x.Check(e)
   for _, each := range []target{
      channel.Pkg.Cargo, channel.Pkg.RustStd, channel.Pkg.Rustc,
   } {
      source := each.Target.X8664PcWindowsGnu.XzUrl
      base := filepath.Base(source)
      cache := filepath.Join(install.Cache, base)
      if ! x.IsFile(cache) {
         _, e := x.Copy(source, cache)
         x.Check(e)
      }
      println(x.ColorCyan("Extract"), base)
      e = tarXz.Unarchive(cache, install.Dest)
      x.Check(e)
   }
}
