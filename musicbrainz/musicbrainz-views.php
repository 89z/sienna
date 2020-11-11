<?php
declare(strict_types = 1);
error_reporting(E_ALL);

require_once 'cove/helper.php';
require_once 'sienna/musicbrainz.php';
require_once 'sienna/youtube.php';

function yt_result(string $query_s): string {
   $query_m['search_query'] = $query_s;
   $res_s = 'https://www.youtube.com/results?' . http_build_query($query_m);
   echo $res_s, "\n";
   $get_s = file_get_contents($res_s);
   preg_match('!/watch\?v=[^"]*!', $get_s, $mat_a);
   return $mat_a[0];
}

if ($argc != 2) {
   echo <<<eof
usage:
musicbrainz-views.php <URL>

examples:
https://musicbrainz.org/release-group/d03bb6b1-d7b4-38ea-974e-847cbb31dca4
https://musicbrainz.org/release/7a629d52-6a61-3ea1-a0a0-dd50bdef63b4
eof;
   exit(1);
}

$url_s = $argv[1];
$mbid_s = basename($url_s);

if (str_contains($url_s, 'release-group')) {
   # RELEASE GROUP
   $releases_a = mb_decode_group($mbid_s);
   $re_n = 0;
   foreach ($releases_a as $n_idx => $o_cur) {
      $re_n = mb_reduce_group($re_n, $o_cur, $n_idx, $releases_a);
   }
   $re_o = $releases_a[$re_n];
} else {
   # RELEASE
   $re_o = mb_decode_release($mbid_s);
}

foreach ($re_o->{'artist-credit'} as $o_artist) {
   $a_out[] = $o_artist->name;
}

$artists_s = implode(' ', $a_out);

foreach ($re_o->media as $media_o) {
   foreach ($media_o->tracks as $track_o) {
      $url_s = yt_result($artists_s . ' ' . $track_o->title);
      $o = new YouTubeViews($url_s);
      echo $o->color(), "\n";
      usleep(500_000);
   }
}
