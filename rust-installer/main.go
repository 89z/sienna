package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "github.com/pelletier/go-toml"
   "os"
   "path"
)

type installer struct {
   cache string
   dest string
}

func newInstaller(dir string) (i installer, e error) {
   cache, e := os.UserCacheDir()
   if e != nil {
      return
   }
   return installer{
      filepath.Join(user, dir),
      filepath.Join(
         filepath.VolumeName(user), dir,
      ),
   }, nil
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

func main() {
   install, e := newInstaller("rust")
   x.Check(e)
   cache := path.Join(
      install.cache, path.Base(source),
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
   e = extract(channel.Pkg.Cargo)
   x.Check(e)
   e = extract(channel.Pkg.RustStd)
   x.Check(e)
   e = extract(channel.Pkg.Rustc)
   x.Check(e)
}
