package main

import (
   "flag"
   "fmt"
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
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
   cmd := exec.Command("ffmpeg", arg...)
   cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
   x.LogInfo("Run", cmd)
   e := cmd.Run()
   if e != nil {
      log.Fatal(e)
   }
}
