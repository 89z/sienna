package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
   "os"
   "path"
)

const (
   remote = "https://static.rust-lang.org/dist/channel-rust-stable.toml"
   local = `C:\rust`
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
         XzUrl string `toml:"xz_url"`
      } `toml:"x86_64-pc-windows-gnu"`
   }
}

type userCache struct {
   dir string
}

func newUserCache(dir ...string) (userCache, error) {
   cache, e := os.UserCacheDir()
   if e != nil {
      return userCache{}, e
   }
   dir = append([]string{cache}, dir...)
   return userCache{
      path.Join(dir...),
   }, nil
}

func (c userCache) install(source string) error {
   base := path.Base(source)
   archive := path.Join(c.dir, base)
   _, e := x.Copy(source, archive)
   if os.IsExist(e) {
      println(x.ColorCyan("Exist"), base)
   } else if e != nil {
      return e
   }
   tar := extract.Tar{Strip: 2}
   println(x.ColorCyan("Xz"), base)
   return tar.Xz(archive, local)
}

func (c userCache) unmarshal(v interface{}) error {
   base := path.Base(remote)
   dest := path.Join(c.dir, base)
   _, e := x.Copy(remote, dest)
   if os.IsExist(e) {
      println(x.ColorCyan("Exist"), base)
   } else if e != nil {
      return e
   }
   data, e := ioutil.ReadFile(dest)
   if e != nil {
      return e
   }
   return toml.Unmarshal(data, v)
}

func main() {
   cache, e := newUserCache("sienna", "rust")
   if e != nil {
      log.Fatal(e)
   }
   var dist distChannel
   e = cache.unmarshal(&dist)
   if e != nil {
      log.Fatal(e)
   }
   for _, each := range []target{
      dist.Pkg.Cargo, dist.Pkg.RustStd, dist.Pkg.Rustc,
   } {
      e = cache.install(each.Target.X8664PcWindowsGnu.XzUrl)
      if e != nil {
         log.Fatal(e)
      }
   }
}
