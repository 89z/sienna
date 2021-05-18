package main

import (
   "fmt"
   "github.com/89z/sienna"
   "net/http"
   "os"
   "path/filepath"
)

const version = "v8.2.2677/gvim_8.2.2677_x64.zip"

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

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "sienna", "vim")
   create := filepath.Join(cache, version)
   _, err = os.Stat(create)
   if err != nil {
      get :=
         "https://github.com/vim/vim-win32-installer/releases/download/" +
         version
      fmt.Println("Get", get)
      r, err := http.Get(get)
      if err != nil {
         panic(err)
      }
      defer r.Body.Close()
      f, err := os.Create(create)
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(r.Body)
   } else {
      fmt.Println("Exist", create)
   }
   arc := sienna.Archive{2}
   fmt.Println("Zip", create)
   arc.Zip(create, `C:\sienna\vim`)
   for _, rt := range runtime {
      get := "https://raw.githubusercontent.com/" + rt.get + rt.create
      fmt.Println("Get", get)
      r, err := http.Get(get)
      if err != nil {
         panic(err)
      }
      defer r.Body.Close()
      f, err := os.Create(filepath.Join(`C:\sienna\vim`, rt.create))
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(r.Body)
   }
}
