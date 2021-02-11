package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "os"
   "path/filepath"
)

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
         Url string
      } `toml:"x86_64-pc-windows-gnu"`
   }
}

const (
   source = "https://static.rust-lang.org/dist/channel-rust-stable.toml"
   dest = `C:\rust`
)

func main() {
   cache, e := os.UserCacheDir()
   x.Check(e)
   cache = filepath.Join(cache, "rust")
   channel := filepath.Join(
      cache, filepath.Base(source),
   )
   _, e = x.Copy(source, channel)
   x.Check(e)
   data, e := ioutil.ReadFile(channel)
   x.Check(e)
   var dist distChannel
   e = toml.Unmarshal(data, &dist)
   x.Check(e)
   tar := extract.Tar{Strip: 2}
   for _, each := range []target{
      dist.Pkg.Cargo, dist.Pkg.RustStd, dist.Pkg.Rustc,
   } {
      source := each.Target.X8664PcWindowsGnu.Url
      archive := filepath.Join(
         cache, filepath.Base(source),
      )
      _, e = x.Copy(source, archive)
      x.Check(e)
      e = tar.Gz(archive, dest)
      x.Check(e)
   }
}
