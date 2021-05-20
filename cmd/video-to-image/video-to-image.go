package main

import (
   "flag"
   "fmt"
   "os"
   "os/exec"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
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
      return
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
   cmd := exec.Command("ffmpeg", arg...)
   cmd.Stderr = os.Stderr
   fmt.Println(invert, "Run", reset, cmd)
   cmd.Run()
}
