<?php
declare(strict_types = 1);
require '/sienna.php';
$s_artist = $_GET['a'];


$s_file = '../../json/' . $_GET[] . ''


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
$s_get = file_get_contents('../../json/' . $s_file . '.json');
$o_get = json_decode($s_get);
$m_local = si_color($o_get->$s_artist);

# MusicBrainz Artist ID
print '<tr><td>@id</td><td>';
if (array_key_exists('q', $_GET)) {
   # write key to file and display
   $o_get->$s_artist->{'@id'} = $_GET['q'];
   $s_json = si_encode($o_get);
} else {
   # prompt for key
}



if (property_exists($o_get->$s_artist, '@id')) {
   print $o_get->$s_artist->{'@id'};
} else {
   print '<form><input name="q"></form>';
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
      printf('<a href="%s">%s</a>', $s_q, $s_album);
   }
   print '</td></tr>';
}
?>
   </table>
   <footer>
      <a href="/add-release?<?= http_build_query($_GET) ?>">Add release</a>
   </footer
</body>
