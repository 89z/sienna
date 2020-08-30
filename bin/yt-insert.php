<?php
declare(strict_types = 1);

require 'sienna/radix-64.php';
require 'sienna/youtube.php';

if ($argc != 2) {
   echo "yt-insert.php <URL>\n";
   exit(1);
}

$s_url = $argv[1];

# year
class YouTubeRelease extends YouTubeInfo {
   function reduce(string $s_ca, string $s_it): string {
      $n_mat = preg_match($s_it, $this->description->simpleText, $a_mat);
      if ($n_mat === 0) {
         return $s_ca;
      }
      $s_mat = $a_mat[1];
      if ($s_mat >= $s_ca) {
         return $s_ca;
      }
      return $s_mat;
   }
}

$a_reg = [
   '/ (\d{4})/',
   '/(\d{4,}) /',
   '/Released on: (\d{4})/',
   '/℗ (\d{4})/'
];

$o_info = new YouTubeRelease($s_url);
$s_init = substr($o_info->publishDate, 0, 4);
$s_year = array_reduce($a_reg, [$o_info, 'reduce'], $s_init);
$n_year = (int)($s_year);

# song, artist
$n_mat = preg_match('/.* · .*/', $o_info->description->simpleText, $a_line);

if ($n_mat !== 0) {
   $s_line = $a_line[0];
   $a_title = explode(' · ', $s_line);
   $a_artist = array_slice($a_title, 1);
   $s_title = implode(', ', $a_artist) . ' - ' . $a_title[0];
} else {
   $s_title = $o_info->title->simpleText;
}

# time
$n_id_1 = time();
$o_rad = new Radix64;
$s_id_1 = $o_rad->encode($n_id_1);

# image
function f_head(string $s_url): bool {
   $a_head = get_headers($s_url);
   $s_code = $a_head[0];
   return strpos($s_code, '200 OK') !== false;
}

if (f_head('https://i.ytimg.com/vi/' . $o_info->id . '/sddefault.jpg')) {
   $s_id_3 = '';
} else if (f_head('https://i.ytimg.com/vi/' . $o_info->id . '/sd1.jpg')) {
   $s_id_3 = '/sd1';
} else {
   var_export($o_info->thumbnail);
   exit(1);
}

# print
$a_rec = [$s_id_1, $n_year, 'y/' . $o_info->id . $s_id_3, $s_title];
$s_json = json_encode($a_rec, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $s_json, ",\n";
