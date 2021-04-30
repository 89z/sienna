package main

import (
   "flag"
   "fmt"
   "github.com/89z/rosso"
   "os"
)

func main() {
   var (
      all bool
      duration, start string
   )
   flag.BoolVar(&all, "a", false, "output all frames")
   flag.StringVar(&duration, "d", "", "duration")
   flag.StringVar(&start, "s", "", "start")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println(`Name:
   video-to-image - create sequence of images from a video

Synopsis:
   video-to-image [flags] <file>

Flags:`)
      flag.PrintDefaults()
      os.Exit(1)
   }
   arg := []string{"-hide_banner"}
   if start != "" {
      arg = append(arg, "-ss", start)
   }
   arg = append(
      arg, "-i", flag.Arg(0),
   )
   if duration != "" {
      arg = append(arg, "-t", duration)
   }
   arg = append(arg, "-q", "1", "-vsync", "vfr")
   if ! all {
      arg = append(arg, "-vf", "select='eq(pict_type, I)'")
   }
   arg = append(arg, "%d.jpg")
   var cmd rosso.Cmd
   err := cmd.Run("ffmpeg", arg...)
   if err != nil {
      panic(err)
   }
}
