package main

import (
   "fmt"
   "log"
   "os"
   "os/exec"
   "path/filepath"
   "strings"
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

func gitCopy(root string) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache = filepath.Join(cache, "sienna", "git")
   dest := root + "git"
   err = os.MkdirAll(
      filepath.Join(dest, "libexec", "git-core"), os.ModeDir,
   )
   if err != nil {
      return err
   }
   err = os.MkdirAll(
      filepath.Join(dest, "share", "git-core", "templates"), os.ModeDir,
   )
   if err != nil {
      return err
   }
   for _, each := range []string{"git.exe", "git-remote-https.exe"} {
      err = os.Link(
         filepath.Join(cache, each),
         filepath.Join(dest, "libexec", "git-core", each),
      )
      if err != nil {
         return err
      }
   }
   return nil
}

func gitMake(root string) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache = filepath.Join(cache, "sienna")
   cmd := exec.Command(
      "git", "clone", "--branch", verGit, "--depth", "1",
      "git://github.com/git/git",
   )
   cmd.Dir = cache
   err = cmd.Run()
   if err != nil {
      cmd = exec.Command("git", "clean", "-d", "-f", "-x")
      cmd.Dir = filepath.Join(cache, "git")
      err = cmd.Run()
      if err != nil {
         return err
      }
   }
   err = os.MkdirAll(
      root + filepath.Join("msys64", "tmp"), os.ModeDir,
   )
   if err != nil {
      return err
   }
   err = os.Setenv("MSYSTEM", "MINGW64")
   if err != nil {
      return err
   }
   paths := []string{
      root + filepath.Join("msys64", "mingw64", "bin"),
      root + filepath.Join("msys64", "usr", "bin"),
   }
   err = os.Setenv("PATH", strings.Join(paths, string(os.PathListSeparator)))
   if err != nil {
      return err
   }
   cmd = exec.Command(
      "make", "-j", "8",
      "CFLAGS=-DCURL_STATICLIB",
      "CURL_LDFLAGS=-lcurl -lwldap32 -lcrypt32",
      "LDFLAGS=-static",
      "NO_GETTEXT=1",
      "NO_ICONV=1",
      "NO_OPENSSL=1",
      "NO_TCLTK=1",
      "USE_LIBPCRE=",
      fmt.Sprintf(`CURLDIR="%v"`, filepath.Join(cache, "curl")),
   )
   cmd.Dir = filepath.Join(cache, "git")
   return cmd.Run()
}

func main() {
   if len(os.Args) != 2 {
      println("git-make <compile | copy>")
      os.Exit(1)
   }
   var (
      err error
      root = os.Getenv("SystemDrive") + "/"
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
