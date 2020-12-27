package main
import "golang.org/x/build/repos"

func Repos() {
   for s, o := range repos.ByImportPath {
      if o.ShowOnDashboard() {
         println(s)
      }
   }
}
