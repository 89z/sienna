package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "github.com/pelletier/go-toml"
   "os"
   "path"
)

const source = "https://static.rust-lang.org/dist/channel-rust-stable.toml"


func unarchive(file, dir string) error {
   tar := &archiver.Tar{OverwriteExisting: true, StripComponents: 2}
   println("EXTRACT", file)
   xz := archiver.TarXz{Tar: tar}
   return xz.Unarchive(file, dir)
}

type distChannel struct{
   Pkg struct{
      Cargo target
      Rustc target
      RustStd target `toml:"rust-std"`
   }
}

type target struct{
   Target struct{
      X86_64PcWindowsGnu struct{
         XzUrl string `toml:"xz_url"`
      } `toml:"x86_64-pc-windows-gnu"`
   }
}

func main() {
   cache, e := x.GetCache("rust")
   x.Check(e)
   dest := path.Join(
      cache, path.Base(source),
   )
   if ! x.IsFile(dest) {
      _, e = x.Copy(source, dest)
      x.Check(e)
   }
   data, e := ioutil.ReadFile(dest)
   x.Check(e)
   channel := new(distChannel)
   e = toml.Unmarshal(data, channel)
   x.Check(e)
   for _, pack := range packages {
      url := manifest.M(pack).S("xz_url")
      base := path.Base(url)
      if ! x.IsFile(base) {
         _, e = x.Copy(url, base)
         x.Check(e)
      }
      e = unarchive(base, `C:\rust`)
      x.Check(e)
   }
}
