package main

func main() {
   r, e := newRepoSearch("1000", "2021-04-12")
   if e != nil {
      panic(e)
   }
   for n := 1; n < 10; n++ {
      r.page(n)
   }
}
