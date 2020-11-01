<?php
declare(strict_types = 1);
error_reporting(E_ALL);

require_once 'sienna/strings.php';
require_once 'sienna/youtube.php';

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
$s_year = $o_info->publishDate;

foreach ($a_reg as $s_reg) {
   $s_year = $o_info->reduce($s_year, $s_reg);
}

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
$s_id_1 = base_convert($n_id_1, 10, 36);

# image
$a_jpg = [
   '/sddefault',
   '/sd1',
   '/hqdefault'
];

foreach ($a_jpg as $s_jpg) {
   $s_url = 'https://i.ytimg.com/vi/' . $o_info->id . $s_jpg . '.jpg';
   echo $s_url, "\n";
   $a_head = get_headers($s_url);
   $s_code = $a_head[0];
   if (str_contains($s_code, '200 OK')) {
      break;
   }
}

if ($s_jpg == '/sddefault') {
   $s_jpg = '';
}

# print
$a_rec = [$s_id_1, $n_year, 'y/' . $o_info->id . $s_jpg, $s_title];
$s_json = json_encode($a_rec, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $s_json, ",\n";
