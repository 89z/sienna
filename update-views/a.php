<?php
declare(strict_types = 1);
error_reporting(E_ALL);


#require_once 'sienna/youtube.php';
require_once 'youtube.php';

chdir('D:\Git\sienna\json');
$o_iter = new FilesystemIterator('.');

foreach ($o_iter as $o_info) {
   $s_name = $o_info->getFilename();
   if ($s_name == 'readme.md') {
      continue;
   }
   $s_get = file_get_contents($s_name);
   $o_json = json_decode($s_get);
   foreach ($o_json as $s_artist => $o_artist) {
      foreach ($o_artist as $s_album => $o_album) {
         if ($s_album == '@check') {
            continue;
         }
         if (! property_exists($o_album, '@id')) {
            continue;
         }
         $s_id = $o_album->{'@id'};
         $o_view = new YouTubeViews('/watch?v=' . $s_id);
         echo $o_view->color(), ' ', $s_artist, ' ', $s_album, "\n";
         usleep(200_000);
      }
   }
}
