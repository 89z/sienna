<?php
declare(strict_types = 1);
extension_loaded('openssl') or die('openssl');

if ($argc != 2) {
   echo "open-graph.php <URL>\n";
   exit(1);
}

$url_s = $argv[1];
$get_s = file_get_contents($url_s);

$re_a = [
   '/ content="([^"]+)" property="og:image"/',
   '/="og:image" content="([^"]+)"/',
   '/="og:video" content="([^"]+)"/',
   '/=og:image content="([^"]+)"/'
];

foreach ($re_a as $re_s) {
   $mat_n = preg_match($re_s, $get_s, $get_a);
   if ($mat_n == 1) {
      echo $get_a[1], "\n";
   }
}
