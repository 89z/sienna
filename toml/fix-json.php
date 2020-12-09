<?php
declare(strict_types = 1);

if ($argc != 3) {
   echo "fix-json.php <in file> <out file>\n";
   exit(1);
}

$in_s = $argv[1];
$out_s = $argv[2];
$get_s = file_get_contents($in_s);
$json_o = json_decode($get_s);

foreach ($json_o as $artist_s => $artist_o) {
   $album_n = 0;
   foreach ($artist_o as $album_s => $album_o) {
      if ($album_s[0] == '@') {
         $artist_m[$artist_s][$album_n][$album_s] = $album_o;
      } else {
         foreach ($album_o as $song_s => $rate_s) {
            $artist_m[$artist_s][$album_n][$album_s][][$song_s] = $rate_s;
         }
      }
      $album_n++;
   }
}

$json_n = JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE;
$json_s = json_encode($artist_m, $json_n);
file_put_contents($out_s, $json_s . "\n");
