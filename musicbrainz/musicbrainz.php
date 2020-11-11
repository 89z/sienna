<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
require_once 'cove/helper.php';

# return release array from group string
function mb_decode_group(string $mbid_s): array {
   # part 1
   $query_m['fmt'] = 'json';
   $query_m['inc'] = 'artist-credits recordings';
   $query_m['release-group'] = $mbid_s;
   $out_s = 'https://musicbrainz.org/ws/2/release?' . http_build_query($query_m);
   # part 2
   $url_r = curl_init($out_s);
   curl_setopt($url_r, CURLOPT_RETURNTRANSFER, true);
   curl_setopt($url_r, CURLOPT_USERAGENT, 'anonymous');
   echo $out_s, "\t";
   # part 3
   $group_s = curl_exec($url_r);
   # part 4
   echo green('OK'), "\n";
   return json_decode($group_s)->releases;
}

# return release object from release string
function mb_decode_release(string $mbid_s): object {
   $query_m['fmt'] = 'json';
   $query_m['inc'] = 'artist-credits recordings';
   $query_s = '?' . http_build_query($query_m);
   $url_s = 'https://musicbrainz.org/ws/2/release/' . $mbid_s . $query_s;
   $url_r = curl_init($url_s);
   echo $url_s, "\n\n";
   curl_setopt($url_r, CURLOPT_USERAGENT, 'anonymous');
   curl_setopt($url_r, CURLOPT_RETURNTRANSFER, true);
   $re_s = curl_exec($url_r);
   return json_decode($re_s);
}

class Release {
   function __construct($release_o) {
      foreach ($release_o as $k => $v) {
         $this->$k = $v;
      }
   }
   function status_b(): bool {
      return $this->status == 'Official';
   }
   function tracks_n(): int {
      $ca_n = 0;
      foreach ($this->media as $it_o) {
         $ca_n += $it_o->{'track-count'};
      }
      return $ca_n;
   }
   function date_b(): bool {
      if (! property_exists($this, 'date')) {
         return false;
      }
      if ($this->date == '') {
         return false;
      }
      return true;
   }
   function date_s(): string {
      return $this->date . '-12-31';
   }
}

# return release object from release array
function mb_reduce_group(
   int $acc_n,
   object $cur_o,
   int $idx_n,
   array $src_a
): int {
   if ($idx_n == 0) {
      return 0;
   }
   $old_o = new Release($src_a[$acc_n]);
   if (! $old_o->date_b()) {
      return $idx_n;
   }
   $new_o = new Release($cur_o);
   if (! $new_o->date_b()) {
      return $acc_n;
   }
   if (! $new_o->status_b()) {
      return $acc_n;
   }
   if ($new_o->date_s() > $old_o->date_s()) {
      return $acc_n;
   }
   if ($new_o->date_s() < $old_o->date_s()) {
      return $idx_n;
   }
   if ($new_o->tracks_n() >= $old_o->tracks_n()) {
      return $acc_n;
   }
   return $idx_n;
}
