package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "net/url"
   "os"
   "path"
)

func main() {
   gvim := url.URL{Scheme: "https", Host: "github.com"}
   gvim.Path = path.Join(
      "vim",
      "vim-win32-installer",
      "releases",
      "download",
      "v8.2.2677",
      "gvim_8.2.2677_x64.zip",
   )
   cache, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }
   inst := newInstall(gvim, cache, `C:\`, "sienna", "vim")
   _, e = x.Copy(inst.source, inst.cache)
   if os.IsExist(e) {
      x.LogInfo("Exist", inst.cache)
   } else if e != nil {
      log.Fatal(e)
   }
   arc := extract.Archive{2}
   x.LogInfo("Zip", inst.cache)
   arc.Zip(inst.cache, inst.dest)
   // PaperColor
   color := url.URL{Scheme: "https", Host: "raw.githubusercontent.com"}
   color.Path = "NLKNguyen/papercolor-theme/e397d18a/colors/PaperColor.vim"
   x.Copy(color.String(), `C:\sienna\vim\colors\PaperColor.vim`)
}
