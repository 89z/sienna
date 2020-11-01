<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('openssl') or die('openssl');
require_once 'cove/helper.php';

class YouTubeInfo {
   function __construct(string $s_watch) {
      # part 1
      $s_query = parse_url($s_watch, PHP_URL_QUERY);
      parse_str($s_query, $m_query);
      # part 2
      $this->id = $m_query['v'];
      # part 3
      $s_info = 'https://www.youtube.com/get_video_info?video_id=' . $this->id;
      echo $s_info, "\n";
      # part 4
      $s_get = file_get_contents($s_info);
      parse_str($s_get, $m_get);
      # part 5
      $s_resp = $m_get['player_response'];
      # part 6
      $o_resp = json_decode($s_resp);
      if (! property_exists($o_resp, 'microformat')) {
         return;
      }
      foreach ($o_resp->microformat->playerMicroformatRenderer as $k => $v) {
         $this->$k = $v;
      }
   }
}

class YouTubeViews extends YouTubeInfo {
   function color(): string {
      if (! property_exists($this, 'viewCount')) {
         return 'undefined property: viewCount';
      }
      $n_views = (int)($this->viewCount);
      $n_then = strtotime($this->publishDate);
      $n_now = time();
      $n_diff = ($n_now - $n_then) / 60 / 60 / 24 / 365;
      $n_rate = $n_views / $n_diff;
      $s_rate = number_format($n_rate);
      if ($n_rate > 8_000_000) {
         return 'RED ' . red($s_rate);
      }
      return 'GREEN ' . green($s_rate);
   }
}
