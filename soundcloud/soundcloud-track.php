<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('openssl') or die('openssl');

if ($argc != 2) {
   echo "soundcloud-track.php <URL>\n";
   exit(1);
}

$UrlS = $argv[1];
$GetS = file_get_contents($UrlS);
$DecS = html_entity_decode($GetS);

# track
preg_match('!/tracks/([^"]*)"!', $DecS, $TrackA);
$AudioS = $TrackA[1];

# img
preg_match('!/artworks-([^.]*)-t500x500\.!', $DecS, $ImgA);
$VideoS = $ImgA[1];

# year
preg_match('/ pubdate>(\d{4})-/', $DecS, $YearA);
$YearS = $YearA[1];
$YearN = (int)($YearS);

# title
preg_match('/<title>([^|]*) by ([^|]*) \|/', $DecS, $TitleA);
$TitleS = $TitleA[2] . ' - ' . $TitleA[1];

# time
$DateN = time();
$DateS = base_convert($DateN, 10, 36);

# print
$RecA = [$DateS, $YearN, 's/' . $AudioS . '/' . $VideoS, $TitleS];
$JsonS = json_encode($RecA, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $JsonS, ",\n";
