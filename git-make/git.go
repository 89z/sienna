package main

import (
   "log"
   "os"
   "os/exec"
   "path/filepath"
)

const (
   verCurl = "curl-7_73_0"
   verGit = "v2.29.1"
)

func curlMake() error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache = filepath.Join(cache, "sienna", "curl")
   err = exec.Command(
      "git", "clone", "--branch", verCurl, "--depth", "1",
      "git://github.com/curl/curl", cache,
   ).Run()
   if err != nil {
      return err
   }
   cmd := exec.Command(
      "mingw32-make", "-f", "Makefile.m32", "-j", "5", "CFG=-winssl",
   )
   cmd.Dir = filepath.Join(cache, "lib")
   return cmd.Run()
}

func gitMake(root string) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cmd := exec.Command(
      "git", "clone", "--branch", verGit, "--depth", "1",
      "git://github.com/git/git",
   )
   cmd.Dir = filepath.Join(cache, "sienna")
   err = cmd.Run()
   git := filepath.Join(cache, "sienna", "git")
   if err != nil {
      cmd = exec.Command("git", "clean", "-d", "-f", "-x")
      cmd.Dir = git
      err = cmd.Run()
      if err != nil {
         return err
      }
   }
   err = os.MkdirAll(`C:\msys64\tmp`, os.ModeDir)
   if err != nil {
      return err
   }
   err = os.Setenv("MSYSTEM", "MINGW64")
   if err != nil {
      return err
   }
   err = os.Setenv("PATH", `C:\msys64\mingw64\bin;C:\msys64\usr\bin`)
   if err != nil {
      return err
   }
   cmd = exec.Command(
      "make", "-j", "8",
      "CFLAGS=-DCURL_STATICLIB",
      // FIXME this needs to be forward slash, or maybe quoted
      "CURLDIR=" + filepath.Join(cache, "sienna", "curl"),
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

func gitCopy(root string) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache = filepath.Join(cache, "sienna", "git")
   core := filepath.Join(root, "git", "libexec", "git-core")
   err = os.MkdirAll(core, os.ModeDir)
   if err != nil {
      return err
   }
   err = os.MkdirAll(
      filepath.Join(root, "git", "share", "git-core", "templates"), os.ModeDir,
   )
   if err != nil {
      return err
   }
   for _, each := range []string{"git.exe", "git-remote-https.exe"} {
      err = os.Link(
         filepath.Join(cache, each), filepath.Join(core, each),
      )
      if err != nil {
         return err
      }
   }
   return nil
}

func main() {
   if len(os.Args) != 2 {
      println("git-make <compile | copy>")
      os.Exit(1)
   }
   var (
      err error
      root = os.Getenv("SystemDrive") + string(os.PathSeparator)
   )
   if os.Args[1] == "copy" {
      err = gitCopy(root)
   } else {
      err = gitMake(root)
   }
   if err != nil {
      log.Fatal(err)
   }
}
