package main

import (
   "flag"
   "fmt"
   "image"
   "image/jpeg"
   "os"
   "path/filepath"
)

func main() {
   var top, right, bottom, left int
   flag.IntVar(&top, "top", 0, "pixels")
   flag.IntVar(&right, "right", 0, "pixels")
   flag.IntVar(&bottom, "bottom", 0, "pixels")
   flag.IntVar(&left, "left", 0, "pixels")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("crop-image [flags] <file>")
      flag.PrintDefaults()
      return
   }
   name := flag.Arg(0)
   in, err := os.Open(name)
   if err != nil {
      panic(err)
   }
   defer in.Close()
   out, err := os.Create("crop-" + filepath.Base(name))
   if err != nil {
      panic(err)
   }
   defer out.Close()
   decode, err := jpeg.Decode(in)
   if err != nil {
      panic(err)
   }
   bound := decode.Bounds()
   rect := image.Rect(left, top, bound.Max.X - right, bound.Max.Y - bottom)
   fmt.Println(bound, rect)
   jpeg.Encode(out, decode.(*image.YCbCr).SubImage(rect), nil)
}
