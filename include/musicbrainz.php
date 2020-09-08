<?php
declare(strict_types = 1);
error_reporting(E_ALL);

extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');

# return release array from group string
function mb_decode_group(string $s_mbid): array {
   $m_q['fmt'] = 'json';
   $m_q['inc'] = 'artist-credits recordings';
   $m_q['release-group'] = $s_mbid;
   $s_out = 'https://musicbrainz.org/ws/2/release?' . http_build_query($m_q);
   $r_c = curl_init($s_out);
   curl_setopt($r_c, CURLOPT_USERAGENT, 'anonymous');
   curl_setopt($r_c, CURLOPT_RETURNTRANSFER, true);
   echo $s_out, "\n";
   $s_group = curl_exec($r_c);
   $o_group = json_decode($s_group);
   return $o_group->releases;
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
   public $s_date;
   public $n_tracks;

   function s_year(): string {
      return substr($this->s_date, 0, 4);
   }

   function n_date_len(): int {
      return strlen($this->s_date);
   }
}

# return release object from release array
function mb_reduce_group(object $o_old, object $o_rel): object {
   if ($o_rel->date == '') {
      return $o_old;
   }
   if ($o_rel->status == 'Promotion') {
      return $o_old;
   }
   $f_sum = fn($n_ca, $o_it) => $n_ca + $o_it->{'track-count'};
   $o_new = new Release;
   $o_new->n_tracks = array_reduce($o_rel->media, $f_sum);
   $o_new->s_date = $o_rel->date;
   if ($o_old == null) {
      return $o_new;
   }
   if ($o_new->s_year() > $o_old->s_year()) {
      return $o_old;
   }
   if ($o_new->s_year() < $o_old->s_year()) {
      return $o_new;
   }
   if ($o_new->n_tracks > $o_old->n_tracks) {
      return $o_old;
   }
   if ($o_new->n_tracks < $o_old->n_tracks) {
      return $o_new;
   }
   if ($o_new->n_date_len() < $o_old->n_date_len()) {
      return $o_old;
   }
   if ($o_new->n_date_len() > $o_old->n_date_len()) {
      return $o_new;
   }
   if ($o_new->s_date >= $o_old->s_date) {
      return $o_old;
   }
   return $o_new;
}
