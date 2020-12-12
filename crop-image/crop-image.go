package main

import (
   "flag"
   "image"
   "image/jpeg"
   "log"
   "os"
   "path/filepath"
)

func main() {
   var bottom_n, left_n, right_n, top_n int
   flag.IntVar(&bottom_n, "bottom", 0, "pixels")
   flag.IntVar(&left_n, "left", 0, "pixels")
   flag.IntVar(&right_n, "right", 0, "pixels")
   flag.IntVar(&top_n, "top", 0, "pixels")

   flag.Parse()
   if flag.NArg() != 1 {
      println("crop-image [flags] <file>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   in_s := flag.Arg(0)
   out_s := "crop-" + filepath.Base(in_s)

   in_file, e := os.Open(in_s)
   if e != nil {
      log.Fatal(e)
   }

   out_file, e := os.Create(out_s)
   if e != nil {
      log.Fatal(e)
   }

   in_img, e := jpeg.Decode(in_file)
   if e != nil {
      log.Fatal(e)
   }

   in_rect := in_img.Bounds()
   out_rect := image.Rect(
      left_n,
      top_n,
      in_rect.Max.X - right_n,
      in_rect.Max.Y - bottom_n,
   )

   out_img := in_img.(*image.YCbCr).SubImage(out_rect)
   jpeg.Encode(out_file, out_img, nil)
}
