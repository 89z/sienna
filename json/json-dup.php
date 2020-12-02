<?php
declare(strict_types = 1);

if ($argc != 3) {
   echo "json-dup.php <artist> <file>\n";
   exit(1);
}

$arg_band_s = $argv[1];
$arg_path_s = $argv[2];
$get_s = file_get_contents($arg_path_s);
$json_o = json_decode($get_s);

foreach ($json_o as $band_s => $band_o) {
   if ($band_s == $arg_band_s) {
      break;
   }
}

$song_m = [];

foreach ($band_o as $album_s => $album_o) {
   if ($album_s[0] == '@') {
      continue;
   }
   foreach ($album_o as $song_s => $rate_s) {
      if ($song_s[0] == '@') {
         continue;
      }
      if (key_exists($song_s, $song_m)) {
         echo $song_s, "\n";
      }
      $song_m[$song_s] = true;
   }
}
