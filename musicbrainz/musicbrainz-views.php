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
$dec_o = new MusicBrainzDecode($mbid_s);

if (str_contains($url_s, 'release-group')) {
   $rel_a = $dec_o->group();
   $rel_n = 0;
   foreach ($rel_a as $idx_n => $cur_o) {
      $rel_n = MusicBrainzReduce($rel_n, $cur_o, $idx_n, $rel_a);
   }
   $rel_o = $rel_a[$rel_n];
} else {
   $rel_o = $dec_o->release();
}

foreach ($rel_o->{'artist-credit'} as $artist_o) {
   $out_a[] = $artist_o->name;
}

$artists_s = implode(' ', $out_a);

foreach ($rel_o->media as $media_o) {
   foreach ($media_o->tracks as $track_o) {
      $url_s = yt_result($artists_s . ' ' . $track_o->title);
      $o = new YouTubeViews($url_s);
      echo $o->color(), "\n";
      usleep(500_000);
   }
}
