<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('openssl') or die('openssl');
require_once 'cove/helper.php';

class YouTubeInfo {
   function __construct(string $WatchS) {
      # part 1
      $QueryS = parse_url($WatchS, PHP_URL_QUERY);
      parse_str($QueryS, $QueryM);
      # part 2
      $this->id = $QueryM['v'];
      # part 3
      $InfoS = 'https://www.youtube.com/get_video_info?video_id=' . $this->id;
      echo $InfoS, "\n";
      # part 4
      $GetS = file_get_contents($InfoS);
      parse_str($GetS, $GetM);
      # part 5
      $RespS = $GetM['player_response'];
      # part 6
      $RespO = json_decode($RespS);
      if (! property_exists($RespO, 'microformat')) {
         return;
      }
      foreach ($RespO->microformat->playerMicroformatRenderer as $k => $v) {
         $this->$k = $v;
      }
   }
}

class YouTubeViews extends YouTubeInfo {
   function color(): string {
      if (! property_exists($this, 'viewCount')) {
         return 'undefined property: viewCount';
      }
      $ViewsN = (int)($this->viewCount);
      $ThenN = strtotime($this->publishDate);
      $NowN = time();
      $DiffN = ($NowN - $ThenN) / 60 / 60 / 24 / 365;
      $RateN = $ViewsN / $DiffN;
      $RateS = number_format($RateN);
      if ($RateN > 8_000_000) {
         return 'RED ' . red($RateS);
      }
      return 'GREEN ' . green($RateS);
   }
}
