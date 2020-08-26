<?php
declare(strict_types = 1);

require 'sienna/sienna.php';

$s_file = $_GET['f'];
$s_artist = $_GET['a'];
$s_rel = $_GET['r'];
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <title><?= $s_rel ?> - Sienna</title>
</head>
<body>
   <header>
      <a href="/artist?<?= http_build_query($_GET) ?>">Up</a>
      <h1><?= "$s_artist - $s_rel" ?></h1>
   </header>
   <table>
<?php
$s_json = '../data/' . $s_file . '.json';
$s_get = file_get_contents($s_json);
$o_get = json_decode($s_get);

if (array_key_exists('s', $_GET)) {
   $s_rate = $_GET['s'];
   $s_track = $_GET['t'];
   $o_get->$s_artist->$s_rel->$s_track = $s_rate;
   $s_enc = si_encode($o_get);
   file_put_contents($s_json, $s_enc . "\n");
}

foreach ($o_get->$s_artist->$s_rel as $s_key => $s_val) {
   printf('<tr><td>%s</td><td>', $s_key);
   if ($s_val == '') {
      print '<form>';
      printf('<input type="hidden" name="f" value="%s">', $s_file);
      printf('<input type="hidden" name="a" value="%s">', $s_artist);
      printf('<input type="hidden" name="r" value="%s">', $s_rel);
      printf('<input type="hidden" name="t" value="%s">', $s_key);
      print '<input name="s">';
      print '</form>';
   } else {
      print $s_val;
   }
   print '</td></tr>';
}
?>
   </table>
</body>
