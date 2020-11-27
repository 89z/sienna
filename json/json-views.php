<?php
declare(strict_types = 1);
require_once 'sienna/youtube.php';

chdir('D:\Git\sienna\json');
$iter_o = new FilesystemIterator('.');

foreach ($iter_o as $info_o) {
   $name_s = $info_o->getFilename();
   if ($name_s == 'readme.md') {
      continue;
   }
   $get_s = file_get_contents($name_s);
   $json_o = json_decode($get_s);
   foreach ($json_o as $artist_s => $artist_o) {
      foreach ($artist_o as $album_s => $album_o) {
         if ($album_s == '@check') {
            continue;
         }
         if (! property_exists($album_o, '@yt')) {
            continue;
         }
         $id_s = $album_o->{'@yt'};
         $view_o = new YouTubeViews('/watch?v=' . $id_s);
         echo $view_o->color(), ' ', $artist_s, ' ', $album_s, "\n";
         usleep(200_000);
      }
   }
}
