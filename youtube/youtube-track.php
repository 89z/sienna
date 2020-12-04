<?php
declare(strict_types = 1);

require 'cove/helper.php';
require 'sienna/youtube.php';

if ($argc != 2) {
   echo "youtube-track.php <URL>\n";
   exit(1);
}

$url_s = $argv[1];
$query_s = parse_url($url_s, PHP_URL_QUERY);
parse_str($query_s, $query_m);
$id_s = $query_m['v'];

# year
$info_o = youtube_info($id_s);

if (! property_exists($info_o, 'description')) {
   echo "Clapham Junction\n";
   exit(1);
}

$year_s = $info_o->publishDate;

$reg_a = [
   '/ (\d{4})/',
   '/(\d{4,}) /',
   '/Released on: (\d{4})/',
   '/℗ (\d{4})/'
];

foreach ($reg_a as $reg_s) {
   $mat_n = preg_match($reg_s, $info_o->description->simpleText, $mat_a);
   if ($mat_n === 0) {
      continue;
   }
   $mat_s = $mat_a[1];
   if ($mat_s >= $year_s) {
      continue;
   }
   $year_s = $mat_s;
}

$year_n = (int)($year_s);

# song, artist
$mat_n = preg_match('/.* · .*/', $info_o->description->simpleText, $line_a);

if ($mat_n !== 0) {
   $line_s = $line_a[0];
   $title_a = explode(' · ', $line_s);
   $artist_a = array_slice($title_a, 1);
   $title_s = implode(', ', $artist_a) . ' - ' . $title_a[0];
} else {
   $title_s = $info_o->title->simpleText;
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
   $url_s = 'https://i.ytimg.com/vi/' . $id_s . $jpg_s . '.jpg';
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
$rec_a = [$date_s, $year_n, 'y/' . $id_s . $jpg_s, $title_s];
$json_s = json_encode($rec_a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $json_s, ",\n";
