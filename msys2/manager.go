package main

import (
   "net/url"
   "strings"
)

var mirror = url.URL{Scheme: "http", Host: "repo.msys2.org"}

func baseName(s, char string) string {
   n := strings.IndexAny(s, char)
   if n == -1 {
      return s
   }
   return s[:n]
}

type database struct {
   name map[string]struct {
      depends []string
      filename string
   }
   provides map[string]string
}
