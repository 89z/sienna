<?php
declare(strict_types = 1);

extension_loaded('openssl') or die('openssl');

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
      foreach ($o_resp->microformat->playerMicroformatRenderer as $k => $v) {
         $this->$k = $v;
      }
   }
}
