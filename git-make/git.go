package main

import (
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
   "path/filepath"
)

const (
   local = `C:\git`
   verCurl = "curl-7_73_0"
   verGit = "v2.29.1"
)

type userCache struct {
   dir string
}

func gitCopy() error {
   var (
      cache userCache
      err error
   )
   cache.dir, err = os.UserCacheDir()
   if err != nil {
      return err
   }
   core := filepath.Join(local, "libexec", "git-core")
   err = os.MkdirAll(core, os.ModeDir)
   if err != nil {
      return err
   }
   err = os.MkdirAll(
      filepath.Join(local, "share", "git-core", "templates"), os.ModeDir,
   )
   if err != nil {
      return err
   }
   cache.dir = filepath.Join(cache.dir, "sienna", "git")
   for _, each := range []string{"git.exe", "git-remote-https.exe"} {
      err = os.Link(
         filepath.Join(cache.dir, each), filepath.Join(core, each),
      )
      if err != nil {
         return err
      }
   }
   return nil
}

func isDir(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && fi.IsDir()
}

func curlMake() error {
   var (
      cache userCache
      err error
   )
   cache.dir, err = os.UserCacheDir()
   if err != nil {
      return err
   }
   cache.dir = filepath.Join(cache.dir, "sienna", "curl")
   err = exec.Command(
      "git", "clone", "--branch", verCurl, "--depth", "1",
      "git://github.com/curl/curl", cache.dir,
   ).Run()
   if err != nil {
      return err
   }
   cmd := exec.Command(
      "mingw32-make", "-f", "Makefile.m32", "-j", "5", "CFG=-winssl",
   )
   cmd.Dir = filepath.Join(cache.dir, "lib")
   return cmd.Run()
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
