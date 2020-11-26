<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('openssl') or die('openssl');

if ($argc != 2) {
   echo "soundcloud-track.php <URL>\n";
   exit(1);
}

$url_s = $argv[1];
$get_s = file_get_contents($url_s);
$dec_s = html_entity_decode($get_s);

# track
preg_match('!/tracks/([^"]*)"!', $dec_s, $TrackA);
$audio_s = $TrackA[1];

# img
preg_match('!/artworks-([^.]*)-t500x500\.!', $dec_s, $ImgA);
$video_s = $ImgA[1];

# year
preg_match('/ pubdate>(\d{4})-/', $dec_s, $YearA);
$year_s = $YearA[1];
$year_n = (int)($year_s);

# title
preg_match('/<title>([^|]*) by ([^|]*) \|/', $dec_s, $title_a);
$title_s = $title_a[2] . ' - ' . $title_a[1];

# time
function encode36(int $n): string {
   $s = (string) $n;
   return base_convert($s, 10, 36);
}

$date_n = time();
$date_s = encode36($date_n);

# print
$rec_a = [$date_s, $year_n, 's/' . $audio_s . '/' . $video_s, $title_s];
$json_s = json_encode($rec_a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $json_s, ",\n";
