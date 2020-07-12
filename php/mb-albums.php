<?php
declare(strict_types = 1);

extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');

function mb_artist($s_arid) {
   # part 1
   $s_url = 'https://musicbrainz.org/ws/2/artist/' . $s_arid . '?fmt=json';
   # part 2
   $r_c = curl_init($s_url);
   curl_setopt($r_c, CURLOPT_RETURNTRANSFER, true);
   curl_setopt($r_c, CURLOPT_USERAGENT, 'anonymous');
   echo $s_url, "\n";
   # part 3
   $s_json = curl_exec($r_c);
   # part 4
   $o_artist = json_decode($s_json);
   return $o_artist->name;
}

function mb_albums($s_arid) {
   $m_q['artist'] = $s_arid;
   $m_q['fmt'] = 'json';
   $m_q['inc'] = 'release-groups';
   $m_q['limit'] = 100;
   $m_q['offset'] = 0;
   $m_q['status'] = 'official';
   $m_q['type'] = 'album';
   $m_remote = [];
   $r_c = curl_init();
   curl_setopt($r_c, CURLOPT_RETURNTRANSFER, true);
   curl_setopt($r_c, CURLOPT_USERAGENT, 'anonymous');
   while (true) {
      # part 1
      $s_q = http_build_query($m_q);
      $s_url = 'https://musicbrainz.org/ws/2/release?' . $s_q;
      curl_setopt($r_c, CURLOPT_URL, $s_url);
      echo $s_url, "\n";
      # part 2
      $s_json = curl_exec($r_c);
      # part 3
      $o_remote = json_decode($s_json);
      foreach ($o_remote->releases as $o_re) {
         $o_rg = $o_re->{'release-group'};
         $a_sec = $o_rg->{'secondary-types'};
         if (count($a_sec) > 0) {
            continue;
         }
         if (array_key_exists($o_rg->title, $m_remote)) {
            continue;
         }
         $m_remote[$o_rg->title] = $o_rg->{'first-release-date'};
      }
      $m_q['offset'] += $m_q['limit'];
      if ($m_q['offset'] >= $o_remote->{'release-count'}) {
         break;
      }
   }
   return $m_remote;
}

if ($argc != 3) {
   echo <<<eof
usage:
mb-albums.php <JSON> <URL>

example:
mb-albums.php Ti-Tz.json `
https://musicbrainz.org/artist/c0b2500e-0cef-4130-869d-732b23ed9df5

eof;
   exit(1);
}

$s_local = $argv[1];
$s_remote = $argv[2];
$s_arid = basename($s_remote);

# remote artist
$s_artist = mb_artist($s_arid);

# local albums
$s_json = file_get_contents($s_local);
# part 2
$o_local = json_decode($s_json);
$s_end = "\e[m";
$s_f_green = "\e[1;32m";
$s_f_red = "\e[1;31m";
$s_f_yellow = "\e[33m";

foreach ($o_local->$s_artist as $s_album => $o_album) {
   if ($s_album == '@check') {
      continue;
   }
   foreach ($o_album as $s_track => $s_rate) {
      if ($s_rate == 'good') {
         $m_local[$s_album] = $s_f_green;
         continue 2;
      }
      if ($s_rate == '') {
         $m_local[$s_album] = $s_f_yellow;
      }
   }
   if (! array_key_exists($s_album, $m_local)) {
      $m_local[$s_album] = $s_f_red;
   }
}

# remote albums
$m_remote = mb_albums($s_arid);
arsort($m_remote);

foreach ($m_remote as $s_title => $s_date) {
   if (array_key_exists($s_title, $m_local)) {
      echo $m_local[$s_title];
   }
   printf("%-10s\t%s%s\n", $s_date, $s_title, $s_end);
}
