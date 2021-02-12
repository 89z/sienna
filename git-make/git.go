package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "path/filepath"
)

const (
   verCurl = "curl-7_73_0"
   verGit = "v2.29.1"
)

func isDir(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && fi.IsDir()
}

func curlMake() (e error) {
   curl, e := x.NewInstall("curl")
   if e != nil {
      return
   }
   if isDir(curl.Cache) {
      return
   }
   e = x.Command(
      "git", "clone", "--branch", verCurl, "--depth", "1",
      "git://github.com/curl/curl", curl.Cache,
   ).Run()
   if e != nil {
      return
   }
   return x.Command(
      "mingw32-make",
      "-C", filepath.Join(curl.Cache, "lib"),
      "-f", "Makefile.m32",
      "-j", "5",
      "CFG=-winssl",
   ).Run()
}

func copyFile(source, dest string) (int64, error) {
   open, e := os.Open(source)
   if e != nil {
      return 0, e
   }
   create, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   defer create.Close()
   return create.ReadFrom(open)
}

func gitCopy() (e error) {
   git, e := x.NewInstall("git")
   if e != nil {
      return
   }
   os.MkdirAll(
      filepath.Join(git.Dest, "share", "git-core", "templates"), os.ModeDir,
   )
   core := filepath.Join(git.Dest, "libexec", "git-core")
   os.MkdirAll(core, os.ModeDir)
   for _, each := range []string{"git.exe", "git-remote-https.exe"} {
      _, e = copyFile(
         filepath.Join(git.Cache, each), filepath.Join(core, each),
      )
      if e != nil {
         return
      }
   }
   return
}

func gitMake(curl string) (e error) {
   git, e := x.NewInstall("git")
   if e != nil {
      return
   }
   if isDir(git.Cache) {
      e = x.Command("git", "-C", git.Cache, "clean", "-d", "-f", "-x").Run()
   } else {
      e = x.Command(
         "git", "clone", "--branch", verGit, "--depth", "1",
         "git://github.com/git/git", git.Cache,
      ).Run()
   }
   if e != nil {
      return
   }
   os.Setenv("MSYSTEM", "MINGW64")
   os.Setenv("PATH", `C:\msys64\mingw64\bin;C:\msys64\usr\bin`)
   os.MkdirAll(`C:\msys64\tmp`, os.ModeDir)
   return x.Command(
      "make", "-C", git.Cache, "-j", "8",
      "CFLAGS=-DCURL_STATICLIB",
      "CURLDIR=" + filepath.ToSlash(curl),
      "CURL_LDFLAGS=-lcurl -lwldap32 -lcrypt32",
      "LDFLAGS=-static",
      "NO_GETTEXT=1",
      "NO_ICONV=1",
      "NO_OPENSSL=1",
      "NO_TCLTK=1",
      "USE_LIBPCRE=",
   ).Run()
}

func main() {
   if len(os.Args) != 2 {
      println("git-make <compile | copy>")
      os.Exit(1)
   }
   if os.Args[1] == "copy" {
      e := gitCopy()
      if e != nil {
         log.Fatal(e)
      }
   } else {
      curl, e := x.NewInstall("curl")
      if e != nil {
         log.Fatal(e)
      }
      e = gitMake(curl.Cache)
      if e != nil {
         log.Fatal(e)
      }
   }
}
