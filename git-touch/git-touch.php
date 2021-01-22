<?php
declare(strict_types = 1);

$r = popen('git ls-files', 'r');
$file_n = 0;

while (true) {
   $get_s = fgets($r);
   if (feof($r)) {
      break;
   }
   $trim_s = rtrim($get_s);
   $file_m[$trim_s] = false;
   $file_n++;
}

$r = popen('git log -m -z --name-only --relative --format=%ct .', 'r');

while ($file_n > 0) {
   $get_s = fgets($r);
   $trim_s = rtrim($get_s);
   $name_a = explode("\x0", $trim_s);
   $unix_s = array_pop($name_a);
   foreach ($name_a as $name_s) {
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
