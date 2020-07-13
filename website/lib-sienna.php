<?php
declare(strict_types = 1);

function si_color($o_artist) {
   foreach ($o_artist as $s_album => $o_album) {
      if (strpos($s_album, '@') === 0) {
         continue;
      }
      foreach ($o_album as $s_track => $s_rate) {
         if ($s_rate == 'good') {
            $m_local[$s_album] = 'green';
            continue 2;
         }
         if ($s_rate == '') {
            $m_local[$s_album] = 'yellow';
         }
      }
      if (! array_key_exists($s_album, $m_local)) {
         $m_local[$s_album] = 'red';
      }
   }
   return $m_local;
}
