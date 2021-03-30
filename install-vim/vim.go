package main

import (
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "log"
   "net/url"
   "os"
   "path"
)

var runtime = []struct{dir, base string}{
   {
      "sienna/vim/colors",
      "NLKNguyen/papercolor-theme/e397d18a/colors/PaperColor.vim",
   }, {
      "sienna/vim/syntax",
      "tpope/vim-markdown/564d7436/syntax/markdown.vim",
   }, {
      "sienna/vim/syntax",
      "vim/vim/a942f9ad/runtime/syntax/javascript.vim",
   }, {
      "sienna/vim/syntax",
      "vim/vim/b9c8312e/runtime/syntax/python.vim",
   },
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
   inst := x.NewInstall("sienna/vim", web.Path)
   inst.SetCache()
   _, e := x.Copy(web.String(), inst.Cache)
   if os.IsExist(e) {
      x.LogInfo("Exist", inst.Cache)
   } else if e != nil {
      log.Fatal(e)
   }
   arc := extract.Archive{2}
   x.LogInfo("Zip", inst.Cache)
   arc.Zip(inst.Cache, inst.Dest)
   web.Host = "raw.githubusercontent.com"
   for _, each := range runtime {
      web.Path = each.base
      inst = x.NewInstall(each.dir, each.base)
      os.Remove(inst.Dest)
      _, e = x.Copy(web.String(), inst.Dest)
      if os.IsExist(e) {
         x.LogInfo("Exist", inst.Dest)
      } else if e != nil {
         log.Fatal(e)
      }
   }
}
