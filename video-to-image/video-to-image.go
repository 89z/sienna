package main

import (
   "flag"
   "github.com/89z/sienna"
   "os"
)

var (
   all_b bool
   dur_s string
   start_s string
)

func main() {
   flag.BoolVar(&all_b, "a", false, "output all frames")
   flag.StringVar(&dur_s, "d", "", "duration")
   flag.StringVar(&start_s, "s", "", "start")

   flag.Parse()
   if flag.NArg() != 1 {
println(`Name:
   video-to-image - create sequence of images from a video

Synopsis:
   video-to-image [flags] <file>

Flags:`)
      flag.PrintDefaults()
      os.Exit(1)
   }
   path_s := flag.Arg(0)
   cmd := []string{"-hide_banner"}
   if start_s != "" {
      cmd = append(cmd, "-ss", start_s)
   }
   cmd = append(cmd, "-i", path_s)
   if dur_s != "" {
      cmd = append(cmd, "-t", dur_s)
   }
   cmd = append(cmd, "-q", "1", "-vsync", "vfr")
   if ! all_b {
      cmd = append(cmd, "-vf", "select='eq(pict_type, I)'")
   }
   cmd = append(cmd, "%d.jpg")
   sienna.System(cmd...)
}
