package main

import (
   "fmt"
   "net/url"
)

const vim = "8.2.2668"
var github = url.URL{Scheme: "https:", Host: "github.com"}

func main() {
   github.Path = fmt.Sprintf(
      "vim/vim-win32-installer/releases/download/v%v/gvim_%[1]v_x64.zip", vim,
   )
}
