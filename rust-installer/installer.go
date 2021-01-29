package main
import "github.com/mholt/archiver/v3"
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
