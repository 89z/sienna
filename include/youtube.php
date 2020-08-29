<?php
declare(strict_types = 1);
extension_loaded('openssl') or die('openssl');

# return video_id from URL
function yt_video_id(string $s_url): string {
   $s_q = parse_url($s_url, PHP_URL_QUERY);
   parse_str($s_q, $m_q);
   return $m_q['v'];
}

# return info object from video id
class YouTubeInfo {
   function __construct(string $s_id) {
      # part 1
      $s_url = 'https://www.youtube.com/get_video_info?video_id=' . $s_id;
      echo $s_url, "\n";
      # part 2
      $s_info = file_get_contents($s_url);
      parse_str($s_info, $m_info);
      # part 3
      $s_resp = $m_info['player_response'];
      # part 4
      $o_resp = json_decode($s_resp);
      foreach ($o_resp->microformat->playerMicroformatRenderer as $s_k => $v) {
         $this->$s_k = $v;
      }
   }
}
