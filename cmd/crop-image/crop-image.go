package main

import (
   "flag"
   "fmt"
   "image"
   "image/jpeg"
   "os"
   "path/filepath"
)

func encode(name string, top, right, bottom, left int) error {
   in, err := os.Open(name)
   if err != nil { return err }
   defer in.Close()
   out, err := os.Create("crop-" + filepath.Base(name))
   if err != nil { return err }
   defer out.Close()
   decode, err := jpeg.Decode(in)
   if err != nil { return err }
   bound := decode.Bounds()
   rect := image.Rect(left, top, bound.Max.X - right, bound.Max.Y - bottom)
   fmt.Println(bound, rect)
   return jpeg.Encode(out, decode.(*image.YCbCr).SubImage(rect), nil)
}

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
   err := encode(flag.Arg(0), top, right, bottom, left)
   if err != nil {
      panic(err)
   }
}
