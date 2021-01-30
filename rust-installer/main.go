package main

import (
   "github.com/89z/x"
   "github.com/pelletier/go-toml"
   "io/ioutil"
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

func main() {
   install, e := x.NewInstall("rust")
   x.Check(e)
   cache := path.Join(
      install.Cache, path.Base(source),
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
   e = extract(install, channel.Pkg.Cargo)
   x.Check(e)
   e = extract(install, channel.Pkg.RustStd)
   x.Check(e)
   e = extract(install, channel.Pkg.Rustc)
   x.Check(e)
}
