package main

import (
   "bufio"
   "log"
   "os/exec"
)

func output(command ...string) (*bufio.Scanner, error) {
   name, arg := command[0], command[1:]
   b, e := exec.Command(name, arg...).Output()
   return bufio.NewScanner(bytes.NewReader(b)), e
}

func popString(a *[]string) string {
   n := len(*a)
   s := (*a)[n - 1]
   *a = (*a)[:n - 1]
   return s
}

func main() {
   gitLs, e := output("git", "ls-files")
   if e != nil {
      log.Fatal(e)
   }
   files := 0
   file := map[string]bool{}
   for gitLs.Scan() {
      file[gitLs.Text()] = false
      files++
   }
   gitLog, e := output(
      "git", "log", "-m", "-z", "--name-only", "--relative", "--format=%ct", ".",
   )
   if e != nil {
      log.Fatal(e)
   }
   for files > 0 {
      gitLog.Scan()
      commit := gitLog.Text()
      names := strings.Split(commit, "\x00")
      unix := popString(&names)
      for _, name := range names {
         if (! key_exists($name_s, $file_m)) {
            continue;
         }
         if ($file_m[$name_s]) {
            continue;
         }
         echo $unix_n, "\t", $name_s, "\n";
         touch($name_s, $unix_n);
         $file_m[$name_s] = true;
         $file_n--;
      }
      $unix_n = (int)($unix_s);
   }
}
