package main

import (
   "github.com/pelletier/go-toml"
   "os"
)

type Map map[string]interface{}

func System(command ...string) error {
   name, arg := command[0], command[1:]
   o := exec.Command(name, arg...)
   o.Stdout = os.Stdout
   return o.Run()
}

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate_s := os.Args[1]
   System("cargo", "new", "rust-deps")
   os.Chdir("rust-deps")
   m := Map{
      "dependencies": Map{"regex": ""},
      "package": Map{"name": "rust-deps", "version": "1.0.0"}
   }
   $dep_s = <<<eof
[package]
name = "$top_s"
version = "1.0.0"
[dependencies]
$crate_s = "$ver_s"
eof;
   $top_s = 'rust-deps';
   file_put_contents('Cargo.toml', $dep_s);
   system('cargo generate-lockfile');
   $get_s = file_get_contents('Cargo.lock');
   preg_match_all('/name = "([^"]*)"/', $get_s, $get_a);
   $name_a = $get_a[1];
   $prev_s = '';
   $dep_n = 0;
   foreach ($name_a as $name_s) {
      if ($name_s == $top_s) {
         continue;
      }
      if ($name_s == $crate_s) {
         continue;
      }
      if ($name_s == $prev_s) {
         continue;
      }
      echo $name_s, "\n";
      $prev_s = $name_s;
      $dep_n++;
   }
   echo "\n", $dep_n, " deps\n";
   chdir('..');
   $dot_n = RecursiveDirectoryIterator::SKIP_DOTS;
   $dir_o = new RecursiveDirectoryIterator($top_s, $dot_n);
   $kid_n = RecursiveIteratorIterator::CHILD_FIRST;
   $iter_o = new RecursiveIteratorIterator($dir_o, $kid_n);
   foreach ($iter_o as $path_o) {
      $path_s = $path_o->getPathname();
      if (is_dir($path_s)) {
         rmdir($path_s);
      } else {
         unlink($path_s);
      }
   }
   rmdir($top_s);
}
