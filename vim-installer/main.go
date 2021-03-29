package main

import (
   "net/url"
   "path"
)

func main() {
   gvim := url.Url{Scheme: "https:", Host: "github.com"}
   gvim.Path = path.Join(
      "vim",
      "vim-win32-installer",
      "releases",
      "download",
      "v8.2.2677",
      "gvim_8.2.2677_x64.zip",
   )
}
