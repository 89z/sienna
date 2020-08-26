<?php
declare(strict_types = 1);

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

# return release object from release array
function mb_reduce_group(object $o_ca, object $o_it): object {
   if ($o_it->date == '') {
      return $o_ca;
   }
   if ($o_it->status == 'Promotion') {
      return $o_ca;
   }
   $o_it->year = substr($o_it->date, 0, 4);
   $f_sum = fn($n_ca, $o_it) => $n_ca + $o_it->{'track-count'};
   $o_it->track = array_reduce($o_it->media, $f_sum);
   $o_it->len = strlen($o_it->date);
   if ($o_it->year > $o_ca->year) {
      return $o_ca;
   }
   if ($o_it->year < $o_ca->year) {
      return $o_it;
   }
   if ($o_it->track > $o_ca->track) {
      return $o_ca;
   }
   if ($o_it->track < $o_ca->track) {
      return $o_it;
   }
   if ($o_it->len < $o_ca->len) {
      return $o_ca;
   }
   if ($o_it->len > $o_ca->len) {
      return $o_it;
   }
   if ($o_it->date >= $o_ca->date) {
      return $o_ca;
   }
   return $o_it;
}
