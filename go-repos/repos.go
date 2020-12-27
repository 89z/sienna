
package main

import (
   "fmt"
   "golang.org/x/build/repos"
)


func main() {
   
   
   for s := range repos.ByImportPath {
      
      
      fmt.Println(s)
      
      
   }
   
   
}

