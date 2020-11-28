<?php
declare(strict_types = 1);

require 'cove/helper.php';
require 'sienna/musicbrainz.php';
require 'sienna/youtube.php';

if ($argc != 2) {
   echo "youtube-views.php <URL>\n";
   exit(1);
}

$url_s = $argv[1];
$query_s = parse_url($url_s, PHP_URL_QUERY);
parse_str($query_s, $query_m);
$id_s = $query_m['v'];
$info_o = youtube_info($id_s);
echo youtube_views($info_o), "\n";
