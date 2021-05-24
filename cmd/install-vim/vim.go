package main

import (
   "net/http"
   "os"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[7m"
   version = "v8.2.2677/gvim_8.2.2677_x64.zip"
)

var runtime = []struct{get, create string}{
   {"NLKNguyen/papercolor-theme/e397d18a/", "colors/PaperColor.vim",},
   {"PProvost/vim-ps1/master/", "ftdetect/ps1.vim"},
   {"PProvost/vim-ps1/master/", "syntax/ps1.vim"},
   {"cespare/vim-toml/master/", "syntax/toml.vim"},
   {"dart-lang/dart-vim-plugin/master/", "syntax/dart.vim"},
   {"google/vim-ft-go/master/", "syntax/go.vim"},
   {"tpope/vim-markdown/564d7436/", "syntax/markdown.vim"},
   {"vim/vim/a942f9ad/runtime/", "syntax/javascript.vim"},
   {"vim/vim/b9c8312e/runtime/", "syntax/python.vim"},
   {"zah/nim.vim/master/", "ftdetect/nim.vim"},
   {"zah/nim.vim/master/", "syntax/nim.vim"},
}

func getCreate(get, create string) error {
   _, err := os.Stat(create)
   if err == nil {
      println(invert, "Exist", reset, create)
      return nil
   }
   println(invert, "Get", reset, get)
   res, err := http.Get(get)
   if err != nil { return err }
   defer res.Body.Close()
   file, err := os.Create(create)
   if err != nil { return err }
   defer file.Close()
   {
      _, err := file.ReadFrom(res.Body)
      return err
   }
}
