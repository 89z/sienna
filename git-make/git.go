package main

import (
   "log"
   "os"
   "os/exec"
   "path"
)

const (
   local = `C:\git`
   verCurl = "curl-7_73_0"
   verGit = "v2.29.1"
)

func curlMake() error {
   cache, e := os.UserCacheDir()
   if e != nil {
      return e
   }
   cache = path.Join(cache, "sienna", "curl")
   e = exec.Command(
      "git", "clone", "--branch", verCurl, "--depth", "1",
      "git://github.com/curl/curl", cache,
   ).Run()
   if e != nil {
      return e
   }
   cmd := exec.Command(
      "mingw32-make", "-f", "Makefile.m32", "-j", "5", "CFG=-winssl",
   )
   cmd.Dir = path.Join(cache, "lib")
   return cmd.Run()
}

func gitCopy() error {
   cache, e := os.UserCacheDir()
   if e != nil {
      return e
   }
   core := path.Join(local, "libexec", "git-core")
   e = os.MkdirAll(core, os.ModeDir)
   if e != nil {
      return e
   }
   e = os.MkdirAll(
      path.Join(local, "share", "git-core", "templates"), os.ModeDir,
   )
   if e != nil {
      return e
   }
   cache = path.Join(cache, "sienna", "git")
   for _, each := range []string{"git.exe", "git-remote-https.exe"} {
      e = os.Link(
         path.Join(cache, each), path.Join(core, each),
      )
      if e != nil {
         return e
      }
   }
   return nil
}

func gitMake() error {
   cache, e := os.UserCacheDir()
   if e != nil {
      return e
   }
   cmd := exec.Command(
      "git", "clone", "--branch", verGit, "--depth", "1",
      "git://github.com/git/git",
   )
   cmd.Dir = path.Join(cache, "sienna")
   e = cmd.Run()
   git := path.Join(cache, "sienna", "git")
   if e != nil {
      cmd = exec.Command("git", "clean", "-d", "-f", "-x")
      cmd.Dir = git
      e = cmd.Run()
      if e != nil {
         return e
      }
   }
   e = os.MkdirAll(`C:\msys64\tmp`, os.ModeDir)
   if e != nil {
      return e
   }
   e = os.Setenv("MSYSTEM", "MINGW64")
   if e != nil {
      return e
   }
   e = os.Setenv("PATH", `C:\msys64\mingw64\bin;C:\msys64\usr\bin`)
   if e != nil {
      return e
   }
   cmd = exec.Command(
      "make", "-j", "8",
      "CFLAGS=-DCURL_STATICLIB",
      "CURLDIR=" + path.Join(cache, "sienna", "curl"),
      "CURL_LDFLAGS=-lcurl -lwldap32 -lcrypt32",
      "LDFLAGS=-static",
      "NO_GETTEXT=1",
      "NO_ICONV=1",
      "NO_OPENSSL=1",
      "NO_TCLTK=1",
      "USE_LIBPCRE=",
   )
   cmd.Dir = git
   return cmd.Run()
}

func main() {
   if len(os.Args) != 2 {
      println("git-make <compile | copy>")
      os.Exit(1)
   }
   var e error
   if os.Args[1] == "copy" {
      e = gitCopy()
   } else {
      e = gitMake()
   }
   if e != nil {
      log.Fatal(e)
   }
}
