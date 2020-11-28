<?php
declare(strict_types = 1);
require_once 'sienna/youtube.php';

if ($argc == 1) {
   echo "json-views.php [-a artist] <file>\n";
   exit(1);
}

$opt_m = getopt('a:', [], $opt_n);
$path_s = $argv[$opt_n];
$get_s = file_get_contents($path_s);
$json_o = json_decode($get_s);

foreach ($json_o as $artist_s => $artist_o) {
   if (key_exists('a', $opt_m)) {
      if ($artist_s != $opt_m['a']) {
         continue;
      }
   }
   foreach ($artist_o as $album_s => $album_o) {
      if ($album_s == '@check') {
         continue;
      }
      if (! property_exists($album_o, '@yt')) {
         continue;
      }
      $id_s = $album_o->{'@yt'};
      $info_o = youtube_info($id_s);
      echo youtube_views($info_o), ' ', $artist_s, ' ', $album_s, "\n";
      usleep(400_000);
   }
}
