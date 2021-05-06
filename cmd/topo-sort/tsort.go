package main

import (
   "bufio"
   "bytes"
   "fmt"
   "os"
   "os/exec"
)

func tsort(graph map[string][]string, end string) []string {
   var (
      b = make(map[string]bool)
      l []string
      s = []string{end}
   )
   for len(s) > 0 {
      n := s[len(s) - 1]
      b[n] = true
      for _, m := range graph[n] {
         if ! b[m] {
            s = append(s, m)
         }
      }
      if s[len(s) - 1] == n {
         s = s[:len(s) - 1]
         l = append(l, n)
      }
   }
   return l
}

func main() {
   if len(os.Args) == 1 {
      fmt.Println("topo-sort <directory> [function]")
      os.Exit(1)
   }
   c := exec.Command("callgraph", "-format", "digraph", os.Args[1])
   if len(os.Args) == 2 {
      c.Stderr, c.Stdout = os.Stderr, os.Stdout
      c.Run()
      return
   }
   b := new(bytes.Buffer)
   c.Stdout = b
   c.Run()
   s := bufio.NewScanner(b)
   m := make(map[string][]string)
   for s.Scan() {
      var parent, child string
      fmt.Sscan(s.Text(), &parent, &child)
      m[child] = append(m[child], parent)
   }
   for n, s := range tsort(m, fmt.Sprintf("%q", os.Args[2])) {
      fmt.Print(n+1, ". ", s, "\n")
   }
}
