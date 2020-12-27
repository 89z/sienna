package main
import "golang.org/x/build/repos"

func main() {
   for s, o := range repos.ByImportPath {
      if o.ShowOnDashboard() {
         println(s)
      }
   }
}
