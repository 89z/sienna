package main

import (
   "html/template"
   "net/http"
)

type date struct {
   Month int
   Days []int
}

func index(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("index.html")
   if err != nil {
      panic(err)
   }
   d := date{
      12, []int{30, 31},
   }
   t.Execute(w, d)
}

func main() {
   http.HandleFunc("/", index)
   println("localhost")
   new(http.Server).ListenAndServe()
}
