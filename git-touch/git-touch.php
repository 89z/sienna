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

$r = popen('git log -m -z --name-only --relative --format=%n%ct .', 'r');
fgets($r);

while ($file_n > 0) {
   $unix = (int)rtrim(fgets($r));
   $names = explode("\x0", rtrim(fgets($r)));
   foreach ($names as $name) {
      if (! key_exists($name, $file_m)) {
         continue;
      }
      if ($file_m[$name]) {
         continue;
      }
      echo $unix, "\t", $name, "\n";
      touch($name, $unix);
      $file_m[$name] = true;
      $file_n--;
   }
}
