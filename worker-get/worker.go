package youtube

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
)

type worker struct {
   addr string
   err chan error
   jobs []*bytes.Buffer
   step int64
}

func (w worker) work(job int) {
   req, err := http.NewRequest("GET", w.addr, nil)
   if err != nil {
      w.err <- err
      return
   }
   req.Header.Set("Range", fmt.Sprintf("bytes=%v-", int64(job) * w.step))
   res, err := new(http.Client).Do(req)
   if err != nil {
      w.err <- err
      return
   }
   defer res.Body.Close()
   w.jobs[job] = new(bytes.Buffer)
   io.CopyN(w.jobs[job], res.Body, w.step)
   w.err <- nil
}

func workerGet(addr string, workers int) ([]*bytes.Buffer, error) {
   if workers < 1 {
      return nil, fmt.Errorf("workers out of range")
   }
   res, err := http.Head(addr)
   if err != nil { return nil, err }
   w := worker{
      addr,
      make(chan error),
      make([]*bytes.Buffer, workers),
      res.ContentLength / int64(workers) + 1,
   }
   for job := range w.jobs {
      go w.work(job)
   }
   for range w.jobs {
      err := <-w.err
      if err != nil { return nil, err }
   }
   return w.jobs, nil
}
