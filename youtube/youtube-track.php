<?php
declare(strict_types = 1);
error_reporting(E_ALL);

require_once 'cove/helper.php';
require_once 'sienna/youtube.php';

if ($argc != 2) {
   echo "youtube-track.php <URL>\n";
   exit(1);
}

$url_s = $argv[1];

# year
class YouTubeRelease extends YouTubeInfo {
   function reduce(string $ca_s, string $it_s): string {
      $mat_n = preg_match($it_s, $this->description->simpleText, $mat_a);
      if ($mat_n === 0) {
         return $ca_s;
      }
      $mat_s = $mat_a[1];
      if ($mat_s >= $ca_s) {
         return $ca_s;
      }
      return $mat_s;
   }
}

$reg_a = [
   '/ (\d{4})/',
   '/(\d{4,}) /',
   '/Released on: (\d{4})/',
   '/℗ (\d{4})/'
];

$rel_o = new YouTubeRelease($url_s);
$year_s = $rel_o->publishDate;

foreach ($reg_a as $reg_s) {
   $year_s = $rel_o->reduce($year_s, $reg_s);
}

$year_n = (int)($year_s);

# song, artist
$mat_n = preg_match('/.* · .*/', $rel_o->description->simpleText, $line_a);

if ($mat_n !== 0) {
   $line_s = $line_a[0];
   $title_a = explode(' · ', $line_s);
   $artist_a = array_slice($title_a, 1);
   $title_s = implode(', ', $artist_a) . ' - ' . $title_a[0];
} else {
   $title_s = $rel_o->title->simpleText;
}

# time
function encode36(int $n): string {
   $s = (string) $n;
   return base_convert($s, 10, 36);
}

$date_n = time();
$date_s = encode36($date_n);

# image
$jpg_a = [
   '/sddefault',
   '/sd1',
   '/hqdefault'
];

foreach ($jpg_a as $jpg_s) {
   $url_s = 'https://i.ytimg.com/vi/' . $rel_o->id . $jpg_s . '.jpg';
   echo $url_s, "\n";
   $head_a = get_headers($url_s);
   $code_s = $head_a[0];
   if (str_contains($code_s, '200 OK')) {
      break;
   }
}

if ($jpg_s == '/sddefault') {
   $jpg_s = '';
}

# print
$rec_a = [$date_s, $year_n, 'y/' . $rel_o->id . $jpg_s, $title_s];
$json_s = json_encode($rec_a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $json_s, ",\n";
