package main
import "strings"

func BadPath(s string) bool {
   if ! strings.HasPrefix(s, "golang.org/x/") {
      return true
   }
   if strings.Contains(s, "/cmd/") {
      return true
   }
   if strings.Contains(s, "/vendor/") {
      return true
   }
   return false
}

var bad_repo = map[string]bool{
   "golang.org/x/arch": true,
   "golang.org/x/benchmarks": true,
   "golang.org/x/blog": true,
   "golang.org/x/build": true,
   "golang.org/x/crypto": true,
   "golang.org/x/debug": true,
   "golang.org/x/mobile": true,
   "golang.org/x/mod": true,
   "golang.org/x/oauth2": true,
   "golang.org/x/perf": true,
   "golang.org/x/review": true,
   "golang.org/x/sync": true,
   "golang.org/x/talks": true,
   "golang.org/x/term": true,
   "golang.org/x/time": true,
   "golang.org/x/tools": true,
   "golang.org/x/website": true,
}

func Download() error {
   return nil
}
