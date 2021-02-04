package main

import (
   "github.com/89z/x"
   "os"
   "path"
)

const (
   curl = "curl-7_73_0"
   git = "v2.29.1"
)

func isDir(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && fi.IsDir()
}

func main() {
   install, e := x.NewInstall("curl")
   x.Check(e)
   if ! isDir(install.Cache) {
      e = x.System(
         "git", "clone", "--branch", curl, "--depth", "1",
         "git://github.com/curl/curl", install.Cache,
      )
      x.Check(e)
      e = x.System(
         "mingw32-make",
         "-C", path.Join(install.Cache, "lib"),
         "-f", "Makefile.m32",
         "-j", "5",
         "CFG=-winssl",
      )
      x.Check(e)
   }
   // FIXME
   if (Test-Path git) {
      Set-Location git
      git clean -d -f -x
   } else {
      git clone --branch $s_git --depth 1 git://github.com/git/git
      Set-Location git
   }
   $env:MSYSTEM = 'MINGW64'
   $env:PATH = 'C:\msys2\mingw64\bin;C:\msys2\usr\bin'
   mingw32-make -j 8 `
   CFLAGS=-DCURL_STATICLIB `
   CURLDIR=../curl `
   CURL_LDFLAGS='-lcurl -lwldap32 -lcrypt32' `
   LDFLAGS='-s -static' `
   NO_GETTEXT=1 `
   NO_ICONV=1 `
   NO_OPENSSL=1 `
   NO_TCLTK=1 `
   USE_LIBPCRE=
}
