package main

import (
   "flag"
   "github.com/89z/x"
   "image"
   "image/jpeg"
   "os"
   "path/filepath"
)

var bottom, left, right, top int

func main() {
   flag.IntVar(&bottom, "bottom", 0, "pixels")
   flag.IntVar(&left, "left", 0, "pixels")
   flag.IntVar(&right, "right", 0, "pixels")
   flag.IntVar(&top, "top", 0, "pixels")
   flag.Parse()
   if flag.NArg() != 1 {
      println("crop-image [flags] <file>")
      flag.PrintDefaults()
      os.Exit(1)
   }
   arg := flag.Arg(0)
   open, e := os.Open(arg)
   x.Check(e)
   create, e := os.Create(
      "crop-" + filepath.Base(arg),
   )
   x.Check(e)
   decode, e := jpeg.Decode(open)
   x.Check(e)
   bound := decode.Bounds()
   rect := image.Rect(left, top, bound.Max.X - right, bound.Max.Y - bottom)
   e = jpeg.Encode(
      create, decode.(*image.YCbCr).SubImage(rect), nil,
   )
   x.Check(e)
}
