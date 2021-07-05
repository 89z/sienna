package main

import (
   "github.com/89z/visage"
   "os"
   "os/exec"
   "path/filepath"
   "strings"
)

type environ map[string][]string

func (e environ) Add(key, value string) {
   e[key] = append(e[key], value)
}

func (e environ) Encode() []string {
   var enc []string
   for key, val := range e {
      join := strings.Join(val, string(os.PathListSeparator))
      enc = append(enc, key + "=" + join)
   }
   return enc
}

func main() {
   if len(os.Args) == 1 {
      println("visage <command> [args]")
      return
   }
   env := make(environ)
   env.Add("ComSpec", `C:\Windows\System32\cmd.exe`)
   env.Add("PATHEXT", ".exe")
   env.Add("PROCESSOR_ARCHITECTURE", "AMD64")
   env.Add("TMP", `C:\Windows\TEMP`)
   for _, pat := range visage.Patterns {
      matches, err := filepath.Glob(`C:\visage\` + pat)
      if err != nil {
         panic(err)
      }
      if matches == nil {
         panic(pat)
      }
      match := matches[0]
      dir, ext := filepath.Dir(match), filepath.Ext(match)
      key := map[string]string{
         ".EXE": "PATH", ".H": "INCLUDE", ".LIB": "LIB",
      }[strings.ToUpper(ext)]
      env.Add(key, dir)
   }
   cmd := exec.Command(os.Args[1], os.Args[2:]...)
   cmd.Stderr = os.Stderr
   cmd.Stdout = os.Stdout
   cmd.Env = env.Encode()
   err := cmd.Run()
   if err != nil {
      panic(err)
   }
}
