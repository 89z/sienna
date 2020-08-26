<?php
declare(strict_types = 1);

require 'sienna/sienna.php';

$s_artist = $_GET['a'];
$s_file = $_GET['f'];
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <title><?= $s_artist ?> - Sienna</title>
</head>
<body>
   <header>
      <a href="..">Up</a>
      <h1><?= $s_artist ?></h1>
   </header>
   <table>
<?php
$s_json = '../data/' . $s_file . '.json';
$s_get = file_get_contents($s_json);
$o_get = json_decode($s_get);
$m_local = si_color($o_get->$s_artist);

# MusicBrainz Artist ID
print '<tr><td>@id</td><td>';

if (property_exists($o_get->$s_artist, '@id')) {
   printf('<a href="/add-release?%s">', http_build_query($_GET));
   print $o_get->$s_artist->{'@id'};
   print '</a>';
} else if (array_key_exists('id', $_GET)) {
   print $_GET['id'];
   $o_get->$s_artist->{'@id'} = $_GET['id'];
   $s_enc = si_encode($o_get);
   file_put_contents($s_json, $s_enc . "\n");
} else {
   print '<form>';
   printf('<input type="hidden" name="f" value="%s">', $s_file);
   printf('<input type="hidden" name="a" value="%s">', $s_artist);
   print '<input name="id">';
   print '</form>';
}

print '</td></tr>';

foreach ($o_get->$s_artist as $s_album => $o_album) {
   if ($s_album == '@id') {
      continue;
   }
   print '<tr>';
   if ($s_album[0] == '@') {
      printf('<td>%s</td><td>%s', $s_album, $o_album);
   } else {
      $s_date = $o_album->{'@date'};
      $s_class = $m_local[$s_album];
      $_GET['r'] = $s_album;
      $s_q = http_build_query($_GET);
      printf('<td>%s</td><td class="%s">', $s_date, $s_class);
      printf('<a href="/release?%s">%s</a>', $s_q, $s_album);
   }
   print '</td></tr>';
}
?>
   </table>
</body>
