<?php
declare(strict_types = 1);

require 'lib-mb.php';

if ($argc != 2) {
   echo <<<eof
usage:
musicbrainz-release.php <link>

link:
https://musicbrainz.org/release/7cc21f46-16b4-4479-844c-e779572ca834
https://musicbrainz.org/release-group/67898886-90bd-3c37-a407-432e3680e872

eof;
   exit(1);
}

$s_url = $argv[1];
$s_mbid = basename($s_url);

if (strpos($s_url, 'release-group') !== false) {
   # RELEASE GROUP
   $a_releases = mb_decode_group($s_mbid);
   $o_init = new stdClass;
   $o_init->year = '9999';
   $o_re = array_reduce($a_releases, 'mb_reduce_group', $o_init);
} else {
   # RELEASE
   $o_re = mb_decode_release($s_mbid);
}

$n_min = 179.5 * 1000;
$n_max = 15 * 60 * 1000;
$m_album[$o_re->title]['@date'] = $o_re->date;

foreach ($o_re->media as $o_media) {
   foreach ($o_media->tracks as $o_track) {
      $n_len = $o_track->length;
      if ($n_len < $n_min) {
         $s_track = 'short';
      } else if ($n_len > $n_max) {
         $s_track = 'long';
      } else {
         $s_track = '';
      }
      $m_r = &$m_album[$o_re->title];
      if (array_key_exists($o_track->title, $m_r)) {
         $m_r[$o_track->number . '. ' . $o_track->title] = $s_track;
      } else {
         $m_r[$o_track->title] = $s_track;
      }
   }
}

$n_opt = JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE;
$s_rate = json_encode($m_album, $n_opt);
echo $s_rate, "\n";
