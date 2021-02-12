package main

import (
   "flag"
   "image"
   "image/jpeg"
   "io"
   "log"
   "os"
   "path/filepath"
)

type imageRect struct {
   bottom, left, right, top int
}

func (r imageRect) Encode(dst io.Writer, src io.Reader) error {
   decode, e := jpeg.Decode(src)
   if e != nil {
      return e
   }
   bound := decode.Bounds()
   rect := image.Rect(
      r.left, r.top, bound.Max.X - r.right, bound.Max.Y - r.bottom,
   )
   return jpeg.Encode(
      dst, decode.(*image.YCbCr).SubImage(rect), nil,
   )
}

func main() {
   var rect imageRect
   flag.IntVar(&rect.bottom, "bottom", 0, "pixels")
   flag.IntVar(&rect.left, "left", 0, "pixels")
   flag.IntVar(&rect.right, "right", 0, "pixels")
   flag.IntVar(&rect.top, "top", 0, "pixels")
   flag.Parse()
   if flag.NArg() != 1 {
      println("crop-image [flags] <file>")
      flag.PrintDefaults()
      os.Exit(1)
   }
   arg := flag.Arg(0)
   open, e := os.Open(arg)
   if e != nil {
      log.Fatal(e)
   }
   create, e := os.Create(
      "crop-" + filepath.Base(arg),
   )
   if e != nil {
      log.Fatal(e)
   }
   e = rect.Encode(open, create)
   if e != nil {
      log.Fatal(e)
   }
}
