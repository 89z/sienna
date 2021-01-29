package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "github.com/pelletier/go-toml"
   "os"
   "path"
)

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
   e = extract(channel.Pkg.Cargo)
   x.Check(e)
   e = extract(channel.Pkg.RustStd)
   x.Check(e)
   e = extract(channel.Pkg.Rustc)
   x.Check(e)
}

var tarXz = archiver.TarXz{
   Tar: &archiver.Tar{OverwriteExisting: true, StripComponents: 2},
}

func extract(pkg target) error {
   remote := pkg.Target.X8664PcWindowsGnu.XzUrl
   local := path.Base(remote)
   if ! x.IsFile(local) {
      _, e = x.Copy(remote, local)
      if e != nil {
         return e
      }
   }
   println(x.ColorGreen("Extract"), local)
   return tarXz.Unarchive(local, `C:\rust`)
}
