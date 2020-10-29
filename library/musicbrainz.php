<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
require 'cove/color.php';

# return release array from group string
function mb_decode_group(string $s_mbid): array {
   # part 1
   $m_q['fmt'] = 'json';
   $m_q['inc'] = 'artist-credits recordings';
   $m_q['release-group'] = $s_mbid;
   $s_out = 'https://musicbrainz.org/ws/2/release?' . http_build_query($m_q);
   # part 2
   $r_c = curl_init($s_out);
   curl_setopt($r_c, CURLOPT_RETURNTRANSFER, true);
   curl_setopt($r_c, CURLOPT_USERAGENT, 'anonymous');
   echo $s_out, "\t";
   # part 3
   $s_group = curl_exec($r_c);
   # part 4
   $o_co = new Color;
   echo $o_co->green('OK'), "\n";
   return json_decode($s_group)->releases;
}

# return release object from release string
function mb_decode_release(string $s_mbid): object {
   $m_q['fmt'] = 'json';
   $m_q['inc'] = 'artist-credits recordings';
   $s_q = '?' . http_build_query($m_q);
   $s_url = 'https://musicbrainz.org/ws/2/release/' . $s_mbid . $s_q;
   $r_c = curl_init($s_url);
   echo $s_url, "\n\n";
   curl_setopt($r_c, CURLOPT_USERAGENT, 'anonymous');
   curl_setopt($r_c, CURLOPT_RETURNTRANSFER, true);
   $s_re = curl_exec($r_c);
   return json_decode($s_re);
}

class Release {
   function __construct($o_release) {
      foreach ($o_release as $k => $v) {
         $this->$k = $v;
      }
   }
   function b_status(): bool {
      return $this->status == 'Official';
   }
   function n_tracks(): int {
      $n_ca = 0;
      foreach ($this->media as $o_it) {
         $n_ca += $o_it->{'track-count'};
      }
      return $n_ca;
   }
   function b_date(): bool {
      if (! property_exists($this, 'date')) {
         return false;
      }
      if ($this->date == '') {
         return false;
      }
      return true;
   }
   function s_date(): string {
      return $this->date . '-12-31';
   }
}

# return release object from release array
function mb_reduce_group(
   int $n_acc,
   object $o_cur,
   int $n_idx,
   array $a_src
): int {
   if ($n_idx == 0) {
      return 0;
   }
   $o_old = new Release($a_src[$n_acc]);
   if (! $o_old->b_date()) {
      return $n_idx;
   }
   $o_new = new Release($o_cur);
   if (! $o_new->b_date()) {
      return $n_acc;
   }
   if (! $o_new->b_status()) {
      return $n_acc;
   }
   if ($o_new->s_date() > $o_old->s_date()) {
      return $n_acc;
   }
   if ($o_new->s_date() < $o_old->s_date()) {
      return $n_idx;
   }
   if ($o_new->n_tracks() >= $o_old->n_tracks()) {
      return $n_acc;
   }
   return $n_idx;
}
