<?php
declare(strict_types = 1);
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');

if (! key_exists('f', $_GET)) {
   echo "localhost:10/add-release.php?f=H-J.json&a=Harold Budd";
   exit(1);
}

$file_s = $_GET['f'];
$artist_s = $_GET['a'];

# local albums
$s_json = file_get_contents($file_s);
$o_local = json_decode($s_json);
$s_arid = $o_local->$artist_s->{'@mb'};
$m_local = si_color($o_local->$artist_s);

function si_color(object $o_artist): array {
   foreach ($o_artist as $s_album => $o_album) {
      if ($s_album[0] == '@') {
         continue;
      }
      $b_good = false;
      $b_done = true;
      foreach ($o_album as $s_track => $s_rate) {
         if ($s_track == '@id') {
            $m_local[$s_album] = 'black';
            continue 2;
         }
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

# remote albums
function mb_albums(string $s_arid): array {
   $m_q['artist'] = $s_arid;
   $m_q['fmt'] = 'json';
   $m_q['inc'] = 'release-groups';
   $m_q['limit'] = 100;
   $m_q['offset'] = 0;
   $m_q['status'] = 'official';
   $m_q['type'] = 'album';
   $m_remote = [];
   $r_c = curl_init();
   curl_setopt($r_c, CURLOPT_RETURNTRANSFER, true);
   curl_setopt($r_c, CURLOPT_USERAGENT, 'anonymous');
   while (true) {
      # part 1
      $s_q = http_build_query($m_q);
      $s_url = 'https://musicbrainz.org/ws/2/release?' . $s_q;
      curl_setopt($r_c, CURLOPT_URL, $s_url);
      print $s_url . "\n";
      # part 2
      $s_json = curl_exec($r_c);
      # part 3
      $o_remote = json_decode($s_json);
      foreach ($o_remote->releases as $o_re) {
         $o_rg = $o_re->{'release-group'};
         $a_sec = $o_rg->{'secondary-types'};
         if (count($a_sec) > 0) {
            continue;
         }
         if (array_key_exists($o_rg->title, $m_remote)) {
            continue;
         }
         $m_remote[$o_rg->title] = $o_rg->{'first-release-date'};
      }
      $m_q['offset'] += $m_q['limit'];
      if ($m_q['offset'] >= $o_remote->{'release-count'}) {
         break;
      }
   }
   return $m_remote;
}

$m_remote = mb_albums($s_arid);
arsort($m_remote);
?>
<head>
   <link rel="stylesheet" href="/sienna.css">
</head>
<body>
   <header>
      <a href="/artist?<?= http_build_query($_GET) ?>">Up</a>
      <h1><?= $artist_s ?></h1>
   </header>
   <table>
<?php
foreach ($m_remote as $s_title => $s_date) {
   printf('<tr><td>%s</td>', $s_date);
   if (array_key_exists($s_title, $m_local)) {
      $s_class = $m_local[$s_title];
      printf('<td style="background:%s">%s', $s_class, $s_title);
   } else {
      printf('<td>%s', $s_title);
   }
   print '</td></tr>';
}
?>
   </table>
</body>
