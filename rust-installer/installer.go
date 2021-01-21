package main

import (
   "github.com/89z/x"
   "github.com/89z/x/toml"
   "github.com/mholt/archiver/v3"
   "log"
   "os"
   "path/filepath"
)

const channel = "https://static.rust-lang.org/dist/channel-rust-stable.toml"

var packages = []string{
   "pkg.cargo.target.x86_64-pc-windows-gnu",
   "pkg.rust-std.target.x86_64-pc-windows-gnu",
   "pkg.rustc.target.x86_64-pc-windows-gnu",
}

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func unarchive(file, dir string) error {
   tar := &archiver.Tar{OverwriteExisting: true, StripComponents: 2}
   println("EXTRACT", file)
   xz := archiver.TarXz{Tar: tar}
   return xz.Unarchive(file, dir)
}

func main() {
   user, e := os.UserCacheDir()
   check(e)
   e = os.Chdir(filepath.Join(user, "rust"))
   check(e)
   dist := filepath.Base(channel)
   if ! x.IsFile(dist) {
      _, e = x.HttpCopy(channel, dist)
      check(e)
   }
   manifest, e := toml.LoadFile(dist)
   check(e)
   for _, pack := range packages {
      url := manifest.M(pack).S("xz_url")
      base := filepath.Base(url)
      if ! x.IsFile(base) {
         _, e = x.HttpCopy(url, base)
         check(e)
      }
      e = unarchive(base, `C:\rust`)
      check(e)
   }
}
