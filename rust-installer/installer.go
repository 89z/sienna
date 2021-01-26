package main

import (
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "github.com/mholt/archiver/v3"
   "os"
   "path/filepath"
)

const channel = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

var packages = []string{
   "pkg.cargo.target.x86_64-pc-windows-gnu",
   "pkg.rust-std.target.x86_64-pc-windows-gnu",
   "pkg.rustc.target.x86_64-pc-windows-gnu",
}

func unarchive(file, dir string) error {
   tar := &archiver.Tar{OverwriteExisting: true, StripComponents: 2}
   println("EXTRACT", file)
   xz := archiver.TarXz{Tar: tar}
   return xz.Unarchive(file, dir)
}

func main() {
   user, e := os.UserCacheDir()
   x.Check(e)
   e = os.Chdir(filepath.Join(user, "rust"))
   x.Check(e)
   dist := filepath.Base(channel)
   if ! x.IsFile(dist) {
      _, e = x.Copy(channel, dist)
      x.Check(e)
   }
   manifest, e := toml.LoadFile(dist)
   x.Check(e)
   for _, pack := range packages {
      url := manifest.M(pack).S("xz_url")
      base := filepath.Base(url)
      if ! x.IsFile(base) {
         _, e = x.Copy(url, base)
         x.Check(e)
      }
      e = unarchive(base, `C:\rust`)
      x.Check(e)
   }
}
