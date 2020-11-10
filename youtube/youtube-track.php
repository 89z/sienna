<?php
declare(strict_types = 1);
error_reporting(E_ALL);

require_once 'cove/helper.php';
require_once 'sienna/youtube.php';

if ($argc != 2) {
   echo "youtube-track.php <URL>\n";
   exit(1);
}

$UrlS = $argv[1];

# year
class YouTubeRelease extends YouTubeInfo {
   function reduce(string $CarS, string $ItemS): string {
      $MatN = preg_match($ItemS, $this->description->simpleText, $MatA);
      if ($MatN === 0) {
         return $CarS;
      }
      $MatS = $MatA[1];
      if ($MatS >= $CarS) {
         return $CarS;
      }
      return $MatS;
   }
}

$RegA = [
   '/ (\d{4})/',
   '/(\d{4,}) /',
   '/Released on: (\d{4})/',
   '/℗ (\d{4})/'
];

$RelO = new YouTubeRelease($UrlS);
$YearS = $RelO->publishDate;

foreach ($RegA as $RegS) {
   $YearS = $RelO->reduce($YearS, $RegS);
}

$YearN = (int)($YearS);

# song, artist
$MatN = preg_match('/.* · .*/', $RelO->description->simpleText, $LineA);

if ($MatN !== 0) {
   $LineS = $LineA[0];
   $TitleA = explode(' · ', $LineS);
   $ArtistA = array_slice($TitleA, 1);
   $TitleS = implode(', ', $ArtistA) . ' - ' . $TitleA[0];
} else {
   $TitleS = $RelO->title->simpleText;
}

# time
$DateN = time();
$DateS = base_convert($DateN, 10, 36);

# image
$JpgA = [
   '/sddefault',
   '/sd1',
   '/hqdefault'
];

foreach ($JpgA as $JpgS) {
   $UrlS = 'https://i.ytimg.com/vi/' . $RelO->id . $JpgS . '.jpg';
   echo $UrlS, "\n";
   $HeadA = get_headers($UrlS);
   $CodeS = $HeadA[0];
   if (str_contains($CodeS, '200 OK')) {
      break;
   }
}

if ($JpgS == '/sddefault') {
   $JpgS = '';
}

# print
$RecA = [$DateS, $YearN, 'y/' . $RelO->id . $JpgS, $TitleS];
$JsonS = json_encode($RecA, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $JsonS, ",\n";
