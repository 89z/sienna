package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "github.com/pelletier/go-toml"
   "log"
   "os"
   "path"
)

const remote = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

type distChannel struct {
   Pkg struct {
      Cargo target
      RustStd target `toml:"rust-std"`
      Rustc target
   }
}

type target struct {
   Target struct {
      X8664PcWindowsGnu struct {
         XzUrl string `toml:"xz_url"`
      } `toml:"x86_64-pc-windows-gnu"`
   }
}

type userCache struct {
   dir string
}

func (c userCache) install(source string) error {
   base := path.Base(source)
   archive := path.Join(c.dir, base)
   _, err := x.Copy(source, archive)
   if os.IsExist(err) {
      x.LogInfo("Exist", base)
   } else if err != nil {
      return err
   }
   tar := extract.Archive{2}
   x.LogInfo("Xz", base)
   return tar.Xz(
      archive, os.Getenv("SystemDrive") + string(os.PathSeparator) + "rust",
   )
}

func (c userCache) unmarshal(v interface{}) error {
   base := path.Base(remote)
   dest := path.Join(c.dir, base)
   _, err := x.Copy(remote, dest)
   if os.IsExist(err) {
      x.LogInfo("Exist", base)
   } else if err != nil {
      return err
   }
   data, err := os.ReadFile(dest)
   if err != nil {
      return err
   }
   return toml.Unmarshal(data, v)
}
