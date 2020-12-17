package main

func main() {
   $top_s = 'rust-deps';
   if ($argc == 1) {
      echo <<<eof
${top_s}.php [flags] <crate>
-v string
   version
eof;
      exit(1);
   }
   $opt_m = getopt('v:', [], $opt_n);
   $ver_s = key_exists('v', $opt_m) ? $opt_m['v'] : '';
   $crate_s = $argv[$opt_n];
   system('cargo new ' . $top_s);
   chdir($top_s);
   $dep_s = <<<eof
[package]
name = "$top_s"
version = "1.0.0"
[dependencies]
$crate_s = "$ver_s"
eof;
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
