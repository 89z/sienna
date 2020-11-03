<?php
declare(strict_types = 1);
error_reporting(E_ALL);

require_once 'cove/helper.php';
require_once 'sienna/musicbrainz.php';
require_once 'sienna/youtube.php';

if ($argc != 2) {
   echo "youtube-views.php <URL>\n";
   exit(1);
}

$s_url = $argv[1];
$o = new YouTubeViews($s_url);
echo $o->color();
