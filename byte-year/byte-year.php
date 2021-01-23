<?php
declare(strict_types = 1);

if ($argc > 2) {
   echo <<<eof
byte-year.php [-g]
-g   use Git to get file list
eof;
   exit(1);
}

$now_n = time();
$year_sec_n = 365 * 24 * 60 * 60;

if ($argc == 2) {
   if (is_file('config.toml')) {
      exec('git ls-files :!docs', $name_a);
   } else {
      exec('git ls-files :!proposal', $name_a);
   }
} else {
   $file_o = new FilesystemIterator('.');
   foreach ($file_o as $name_o) {
      $name_a[] = $name_o->getFilename();
   }
}

foreach ($name_a as $name_s) {
   $then_n = filemtime($name_s);
   $year_n = ($now_n - $then_n) / $year_sec_n;
   # we are adding a Map to an Array
   $file_a[] = [
      'name' => $name_s,
      'size' => filesize($name_s) * $year_n
   ];
}

$f = fn ($m, $m2) => $m2['size'] <=> $m['size'];
usort($file_a, $f);

foreach ($file_a as $file_m) {
   echo $file_m['size'], ' ', $file_m['name'], "\n";
}
