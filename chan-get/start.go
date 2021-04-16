package main

import (
   "bytes"
   "fmt"
   "io"
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
      addr, res.ContentLength / jobs + 1, make([]*bytes.Buffer, jobs),
   }, nil
}

func (w worker) work(job int, ch chan bool) error {
   req, e := http.NewRequest("GET", w.addr, nil)
   if e != nil { return e }
   req.Header.Set("Range", fmt.Sprintf("bytes=%v-", int64(job) * w.step))
   res, e := new(http.Client).Do(req)
   if e != nil { return e }
   defer res.Body.Close()
   w.jobs[job] = new(bytes.Buffer)
   io.CopyN(w.jobs[job], res.Body, w.step)
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
   dat, e := os.Create("10Mio.dat")
   if e != nil {
      panic(e)
   }
   defer dat.Close()
   for range w.jobs { <-done }
   for _, job := range w.jobs {
      dat.ReadFrom(job)
   }
}
