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
   var i, ss, t string
   flag.StringVar(&i, "i", "", "infile")
   flag.StringVar(&ss, "ss", "", "seek position")
   flag.StringVar(&t, "t", "", "duration")
   flag.Parse()
   if flag.NArg() != 1 {
      println("ffmpeg-firefox [flags] <outfile>")
      flag.PrintDefaults()
      return
   }
   outfile := flag.Arg(0)
   arg := []string{"-hide_banner"}
   if ss != "" {
      arg = append(arg, "-ss", ss)
   }
   if i != "" {
      arg = append(arg, "-i", i)
   }
   if t != "" {
      arg = append(arg, "-t", t)
   }
   arg = append(arg, "-ac", "2")
   arg = append(arg, "-pix_fmt", "yuv420p")
   arg = append(arg, "-vf", "scale=-1:720")
   arg = append(arg, "-y")
   arg = append(arg, outfile)
   cmd := exec.Command("ffmpeg", arg...)
   cmd.Stderr = os.Stderr
   fmt.Println(invert, "Run", reset, cmd)
   err := cmd.Run()
   if err != nil {
      panic(err)
   }
}
