package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "path"
)

const source = "https://static.rust-lang.org/dist/channel-rust-stable.toml"
var channel distChannel

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
   cache := path.Join(
      install.Cache, path.Base(source),
   )
   _, e = x.Copy(source, cache, x.Ignore)
   x.Check(e)
   data, e := ioutil.ReadFile(cache)
   x.Check(e)
   e = toml.Unmarshal(data, &channel)
   x.Check(e)
   for _, each := range []target{
      channel.Pkg.Cargo, channel.Pkg.RustStd, channel.Pkg.Rustc,
   } {
      source := each.Target.X8664PcWindowsGnu.XzUrl
      base := path.Base(source)
      cache := path.Join(install.Cache, base)
      _, e = x.Copy(source, cache, x.Ignore)
      x.Check(e)
      tar := extract.Tar{Strip: 2}
      e = tar.Xz(cache, install.Dest)
      x.Check(e)
   }
}
