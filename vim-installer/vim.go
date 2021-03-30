package main

import (
   "net/url"
   "path/filepath"
)

type install struct {
   source string
   cache string
   dest string
}

func newInstall(source url.URL, cache, dest string, base ...string) install {
   for _, each := range base {
      cache = filepath.Join(cache, each)
      dest = filepath.Join(dest, each)
   }
   src := source.String()
   cache = filepath.Join(cache, filepath.Base(src))
   return install{src, cache, dest}
}
