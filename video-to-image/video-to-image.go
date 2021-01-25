package main

import (
   "flag"
   "github.com/89z/x"
   "os"
)

var (
   all bool
   duration string
   start string
)

func main() {
   flag.BoolVar(&all, "a", false, "output all frames")
   flag.StringVar(&duration, "d", "", "duration")
   flag.StringVar(&start, "s", "", "start")
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

   path := flag.Arg(0)
   arg := []string{"-hide_banner"}
   if start != "" {
      arg = append(arg, "-ss", start)
   }
   arg = append(arg, "-i", path)
   if duration != "" {
      arg = append(arg, "-t", duration)
   }
   arg = append(arg, "-q", "1", "-vsync", "vfr")
   if ! all {
      arg = append(arg, "-vf", "select='eq(pict_type, I)'")
   }
   arg = append(arg, "%d.jpg")
   x.System("ffmpeg", arg...)
}
