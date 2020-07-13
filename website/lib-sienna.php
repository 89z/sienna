<?php
declare(strict_types = 1);

function si_color($o_artist) {
   foreach ($o_artist as $s_album => $o_album) {
      if (strpos($s_album, '@') === 0) {
         continue;
      }
      $b_good = false;
      $b_done = true;
      foreach ($o_album as $s_track => $s_rate) {
         if ($s_rate == 'good') {
            $b_good = true;
         }
         if ($s_rate == '') {
            $b_done = false;
         }
      }
      if ($b_good && $b_done) {
         $m_local[$s_album] = 'green';
      }
      if ($b_good && ! $b_done) {
         $m_local[$s_album] = 'lightgreen';
      }
      if (! $b_good && $b_done) {
         $m_local[$s_album] = 'red';
      }
      if (! $b_good && ! $b_done) {
         $m_local[$s_album] = 'lightred';
      }
   }
   return $m_local;
}
