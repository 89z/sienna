package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
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

type userCache struct {
   dir string
}

func curlMake() error {
   curl, err := x.NewInstall("curl")
   if err != nil {
      return err
   }
   if isDir(curl.Cache) {
      return err
   }
   err = exec.Command(
      "git", "clone", "--branch", verCurl, "--depth", "1",
      "git://github.com/curl/curl", curl.Cache,
   ).Run()
   if err != nil {
      return err
   }
   return exec.Command(
      "mingw32-make",
      "-C", filepath.Join(curl.Cache, "lib"),
      "-f", "Makefile.m32",
      "-j", "5",
      "CFG=-winssl",
   ).Run()
}

func copyFile(source, dest string) (int64, error) {
   open, err := os.Open(source)
   if err != nil {
      return 0, err
   }
   create, err := os.Create(dest)
   if err != nil {
      return 0, err
   }
   defer create.Close()
   return create.ReadFrom(open)
}

func gitCopy() error {
   git, err := x.NewInstall("git")
   if err != nil {
      return err
   }
   os.MkdirAll(
      filepath.Join(git.Dest, "share", "git-core", "templates"), os.ModeDir,
   )
   core := filepath.Join(git.Dest, "libexec", "git-core")
   os.MkdirAll(core, os.ModeDir)
   for _, each := range []string{"git.exe", "git-remote-https.exe"} {
      _, err = copyFile(
         filepath.Join(git.Cache, each), filepath.Join(core, each),
      )
      if err != nil {
         return err
      }
   }
   return nil
}

func gitMake(curl string) error {
   git, err := x.NewInstall("git")
   if err != nil {
      return err
   }
   if isDir(git.Cache) {
      err = exec.Command("git", "-C", git.Cache, "clean", "-d", "-f", "-x").Run()
   } else {
      err = exec.Command(
         "git", "clone", "--branch", verGit, "--depth", "1",
         "git://github.com/git/git", git.Cache,
      ).Run()
   }
   if err != nil {
      return err
   }
   os.Setenv("MSYSTEM", "MINGW64")
   os.Setenv("PATH", `C:\msys64\mingw64\bin;C:\msys64\usr\bin`)
   os.MkdirAll(`C:\msys64\tmp`, os.ModeDir)
   return exec.Command(
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
      err := gitCopy()
      if err != nil {
         log.Fatal(err)
      }
   } else {
      curl, err := x.NewInstall("curl")
      if err != nil {
         log.Fatal(err)
      }
      err = gitMake(curl.Cache)
      if err != nil {
         log.Fatal(err)
      }
   }
}
