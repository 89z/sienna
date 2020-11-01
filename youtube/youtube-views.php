<?php
declare(strict_types = 1);
error_reporting(E_ALL);

require_once 'cove/helper.php';
require_once 'sienna/musicbrainz.php';
require_once 'sienna/youtube.php';

# return artists string from release object
function yt_encode_artists(object $o_in): string {
   foreach ($o_in->{'artist-credit'} as $o_artist) {
      $a_out[] = $o_artist->name;
   }
   return implode(' ', $a_out);
}

# return video_id from search string
function yt_result(string $s_query): string {
   $m_query['search_query'] = $s_query;
   $s_res = 'https://www.youtube.com/results?' . http_build_query($m_query);
   echo $s_res, "\n";
   $s_get = file_get_contents($s_res);
   preg_match('!/watch\?v=[^"]*!', $s_get, $a_mat);
   return $a_mat[0];
}

if ($argc != 2) {
   echo "youtube-views.php <URL>\n";
   exit(1);
}

$s_url = $argv[1];
$o = new YouTubeViews($s_url);
echo $o->color();
