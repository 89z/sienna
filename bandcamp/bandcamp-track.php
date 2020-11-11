<?php
declare(strict_types = 1);
error_reporting(E_ALL);
extension_loaded('openssl') or die('openssl');

if ($argc != 2) {
   echo "bandcamp-track.php <URL>\n";
   exit(1);
}

$url_s = $argv[1];
$info_s = file_get_contents($url_s);

# track
preg_match('!/track=([^/]*)/!', $info_s, $track_a);
$audio_s = $track_a[1];

# img
preg_match('!/img/([^_]*)_!', $info_s, $img_a);
$video_s = $img_a[1];

# year
preg_match("/ (\d{4})\n/", $info_s, $year_a);
$year_s = $year_a[1];
$year_n = (int)($year_s);

# title
preg_match('!<title>(.*) \| (.*)</title>!', $info_s, $title_a);
$title_s = $title_a[2] . ' - ' . $title_a[1];

# time
$date_n = time();
$date_s = base_convert($date_n, 10, 36);

# print
$rec_a = [$date_s, $year_n, 'b/' . $audio_s . '/' . $video_s, $title_s];
$json_s = json_encode($rec_a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $json_s, ",\n";
