package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "net/url"
   "os"
   "path"
)

type install struct { cache, dest string }

func newInstall(dir, base string) install {
   cache := path.Join(dir, path.Base(base))
   return install{
      cache, os.Getenv("SystemDrive") + string(os.PathSeparator) + cache,
   }
}

func (i *install) setCache() error {
   cache, e := os.UserCacheDir()
   if e != nil { return e }
   i.cache, i.dest = path.Join(cache, i.cache), path.Dir(i.dest)
   return nil
}

func main() {
   web := url.URL{Scheme: "https", Host: "github.com"}
   web.Path = path.Join(
      "vim",
      "vim-win32-installer",
      "releases",
      "download",
      "v8.2.2677",
      "gvim_8.2.2677_x64.zip",
   )
   inst := newInstall("sienna/vim", web.Path)
   inst.setCache()
   _, e := x.Copy(web.String(), inst.cache)
   if os.IsExist(e) {
      x.LogInfo("Exist", inst.cache)
   } else if e != nil {
      log.Fatal(e)
   }
   arc := extract.Archive{2}
   x.LogInfo("Zip", inst.cache)
   arc.Zip(inst.cache, inst.dest)
   // PaperColor
   web.Host = "raw.githubusercontent.com"
   web.Path = "NLKNguyen/papercolor-theme/e397d18a/colors/PaperColor.vim"
   inst = newInstall("sienna/vim/colors", web.Path)
   x.Copy(web.String(), inst.dest)
}
