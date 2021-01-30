package main

import (
   "github.com/89z/x"
   "github.com/mholt/archiver/v3"
   "path/filepath"
)

var tarXz = archiver.TarXz{
   Tar: &archiver.Tar{OverwriteExisting: true, StripComponents: 2},
}

func extract(install x.Install, pkg target) error {
   source := pkg.Target.X8664PcWindowsGnu.XzUrl
   base := filepath.Base(source)
   cache := filepath.Join(install.Cache, base)
   if ! x.IsFile(cache) {
      _, e := x.Copy(source, cache)
      if e != nil {
         return e
      }
   }
   println(x.ColorCyan("Extract"), base)
   return tarXz.Unarchive(cache, install.Dest)
}
