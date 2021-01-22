package main

import (
   "log"
   "os/exec"
)

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func system(command ...string) error {
   name, arg := command[0], command[1:]
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   return c.Run()
}

func main() {
   e := system("git", "commit", "--verbose")
   check(e)
   if (is_file('config.toml')) {
      $dir_o = new RecursiveDirectoryIterator('docs');
      $iter_o = new RecursiveIteratorIterator($dir_o);
      echo "UNLINK\n";
      foreach ($iter_o as $info_o) {
         if ($info_o->isFile()) {
            unlink($info_o->getPathname());
         }
      }
      echo "HUGO\n";
      system('hugo');
      echo "GIT ADD\n";
      system('git add .');
      echo "GIT COMMIT\n";
      system('git commit --amend');
   }
   echo "git push\n";
}
