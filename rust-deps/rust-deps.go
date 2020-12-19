package main

import (
   "log"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate_s := os.Args[1]
   e := System("cargo", "new", "rust-deps")
   if e != nil {
      log.Fatal(e)
   }
   os.Chdir("rust-deps")
   toml_m := Map{
      "dependencies": Map{crate_s: ""},
      "package": Map{"name": "rust-deps", "version": "1.0.0"},
   }
   e = TomlEncode("Cargo.toml", toml_m)
   if e != nil {
      log.Fatal(e)
   }
   e = System("cargo", "generate-lockfile")
   if e != nil {
      log.Fatal(e)
   }
   lock_m, e := TomlDecode("Cargo.lock")
   if e != nil {
      log.Fatal(e)
   }
   pac_a := lock_m.A("package")
   for n := range pac_a {
      name_s := pac_a.M(n).S("name")
      println(name_s)
   }
   /*
   preg_match_all('/name = "([^"]*)"/', $get_s, $get_a);
   $name_a = $get_a[1];
   $prev_s = '';
   $dep_n = 0;
   $top_s = 'rust-deps';
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
   */
}
