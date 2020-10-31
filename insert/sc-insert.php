<?php
declare(strict_types = 1);
error_reporting(E_ALL);
extension_loaded('openssl') or die('openssl');

if ($argc != 2) {
   echo "sc-insert.php <URL>\n";
   exit(1);
}

$s_url = $argv[1];
$s_get = file_get_contents($s_url);
$s_dec = html_entity_decode($s_get);

# track
preg_match('!/tracks/([^"]*)"!', $s_dec, $a_track);
$s_id_2 = $a_track[1];

# img
preg_match('!/artworks-([^.]*)-t500x500\.!', $s_dec, $a_img);
$s_id_3 = $a_img[1];

# year
preg_match('/ pubdate>(\d{4})-/', $s_dec, $a_year);
$s_year = $a_year[1];
$n_year = (int)($s_year);

# title
preg_match('/<title>([^|]*) by ([^|]*) \|/', $s_dec, $a_title);
$s_title = $a_title[2] . ' - ' . $a_title[1];

# time
$n_id_1 = time();
$s_id_1 = base_convert($n_id_1, 10, 36);

# print
$a_rec = [$s_id_1, $n_year, 's/' . $s_id_2 . '/' . $s_id_3, $s_title];
$s_json = json_encode($a_rec, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
echo $s_json, ",\n";
