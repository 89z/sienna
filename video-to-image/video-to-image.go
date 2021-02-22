package main

import (
   "flag"
   "fmt"
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
)

type flags struct {
   all bool
   duration string
   start string
}

func main() {
   var f flags
   flag.BoolVar(&f.all, "a", false, "output all frames")
   flag.StringVar(&f.duration, "d", "", "duration")
   flag.StringVar(&f.start, "s", "", "start")
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
   if f.start != "" {
      arg = append(arg, "-ss", f.start)
   }
   arg = append(
      arg, "-i", flag.Arg(0),
   )
   if f.duration != "" {
      arg = append(arg, "-t", f.duration)
   }
   arg = append(arg, "-q", "1", "-vsync", "vfr")
   if ! f.all {
      arg = append(arg, "-vf", "select='eq(pict_type, I)'")
   }
   arg = append(arg, "%d.jpg")
   cmd := exec.Command("ffmpeg", arg...)
   cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
   fmt.Println(x.ColorCyan("Run"), "ffmpeg", arg)
   e := cmd.Run()
   if e != nil {
      log.Fatal(e)
   }
}
