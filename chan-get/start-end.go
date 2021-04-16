package main

import (
   "bytes"
   "fmt"
   "net/http"
   "os"
)

type worker struct {
   addr string
   step int64
   jobs []*bytes.Buffer
}

func newWorker(addr string, jobs int64) (worker, error) {
   res, e := http.Head(addr)
   if e != nil {
      return worker{}, e
   }
   return worker{
      addr, res.ContentLength / jobs, make([]*bytes.Buffer, jobs),
   }, nil
}

func (w worker) work(job int, ch chan bool) error {
   req, e := http.NewRequest("GET", w.addr, nil)
   if e != nil { return e }
   pos := int64(job) * (w.step + 1)
   req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", pos, pos + w.step))
   res, e := new(http.Client).Do(req)
   if e != nil { return e }
   defer res.Body.Close()
   w.jobs[job] = new(bytes.Buffer)
   w.jobs[job].ReadFrom(res.Body)
   println("END", job)
   ch <- true
   return nil
}

func main() {
   done := make(chan bool)
   w, e := newWorker("http://speedtest.lax.hivelocity.net/10Mio.dat", 2)
   if e != nil {
      panic(e)
   }
   for job := range w.jobs {
      println("BEGIN", job)
      go w.work(job, done)
   }
   bad, e := os.Create("bad.dat")
   if e != nil {
      panic(e)
   }
   defer bad.Close()
   for range w.jobs { <-done }
   for _, job := range w.jobs {
      bad.ReadFrom(job)
   }
}
