<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');

class MusicBrainzDecode {
   var $api_s = 'https://musicbrainz.org/ws/2/release';
   var $query_m = ['fmt' => 'json', 'inc' => 'artist-credits recordings'];

   function __construct(string $mbid_s) {
      $this->mbid_s = $mbid_s;
      $this->url_r = curl_init();
      curl_setopt($this->url_r, CURLOPT_RETURNTRANSFER, true);
      curl_setopt($this->url_r, CURLOPT_USERAGENT, 'anonymous');
   }

   function group(): array {
      $this->query_m['release-group'] = $this->mbid_s;
      $url_s = $this->api_s . '?' . http_build_query($this->query_m);
      curl_setopt($this->url_r, CURLOPT_URL, $url_s);
      echo $url_s, "\n";
      $json_s = curl_exec($this->url_r);
      return json_decode($json_s)->releases;
   }

   function release(): object {
      $query_s = http_build_query($this->query_m);
      $url_s = $this->api_s . '/' . $this->mbid_s . '?' . $query_s;
      curl_setopt($this->url_r, CURLOPT_URL, $url_s);
      echo $url_s, "\n";
      $json_s = curl_exec($this->url_r);
      return json_decode($json_s);
   }
}

class MusicBrainzRelease {
   function __construct($release_o) {
      foreach ($release_o as $k => $v) {
         $this->$k = $v;
      }
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

   function status(): bool {
      return $this->status == 'Official';
   }

   function tracks(): int {
      $ca_n = 0;
      foreach ($this->media as $it_o) {
         $ca_n += $it_o->{'track-count'};
      }
      return $ca_n;
   }
}

function MusicBrainzReduce(
   int $acc_n, object $cur_o, int $idx_n, array $src_a
): int {
   if ($idx_n == 0) {
      return 0;
   }
   $old_o = new MusicBrainzRelease($src_a[$acc_n]);
   if (! $old_o->date_b()) {
      return $idx_n;
   }
   $new_o = new MusicBrainzRelease($cur_o);
   if (! $new_o->date_b()) {
      return $acc_n;
   }
   if (! $new_o->status()) {
      return $acc_n;
   }
   if ($new_o->date_s() > $old_o->date_s()) {
      return $acc_n;
   }
   if ($new_o->date_s() < $old_o->date_s()) {
      return $idx_n;
   }
   if ($new_o->tracks() >= $old_o->tracks()) {
      return $acc_n;
   }
   return $idx_n;
}
