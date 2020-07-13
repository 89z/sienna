<?php
declare(strict_types = 1);
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$s_local = $_GET['f'];
$s_artist = $_GET['a'];

# local albums
$s_json = file_get_contents('../json/' . $s_local);
$o_local = json_decode($s_json);
$s_arid = $o_local->$s_artist->{'@id'};

foreach ($o_local->$s_artist as $s_album => $o_album) {
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

# remote albums
function mb_albums($s_arid) {
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
      echo $s_url, "\n";
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
   <table>
<?php
foreach ($m_remote as $s_title => $s_date) {
   echo '<tr><td>' . $s_date . '</td>';
   if (array_key_exists($s_title, $m_local)) {
      $s_class = $m_local[$s_title];
      echo <<<eof
<td class="$s_class">$s_title</td>
eof;
   } else {
      echo '<td>' . $s_title . '</td>';
   }
   echo '</tr>';
}
?>
   </table>
</body>
