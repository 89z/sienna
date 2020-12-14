package main

import (
   "flag"
   "fmt"
   "os"
   "os/exec"
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

   cmd_a := []string{"-hide_banner"}
   if start_s != "" {
      cmd_a = append(cmd_a, "-ss", start_s)
   }

   cmd_a = append(cmd_a, "-i", path_s)
   if dur_s != "" {
      cmd_a = append(cmd_a, "-t", dur_s)
   }

   cmd_a = append(cmd_a, "-q", "1", "-vsync", "vfr")
   if ! all_b {
      cmd_a = append(cmd_a, "-vf", "select='eq(pict_type, I)'")
   }

   cmd_a = append(cmd_a, "%d.jpg")
   fmt.Printf("ffmpeg %q\n", cmd_a)
   o := exec.Command("ffmpeg", cmd_a...)
   o.Stderr = os.Stderr
   o.Run()
}
