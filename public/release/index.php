<?php
declare(strict_types = 1);
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
      <a href="/artist=?<?= http_build_query($_GET) ?>">Up</a>
      <h1><?= "$s_artist - $s_rel" ?></h1>
   </header>
   <table>
<?php
$s_get = file_get_contents('../json/' . $s_file . '.json');
$o_get = json_decode($s_get);

foreach ($o_get->$s_artist->$s_rel as $s_key => $s_val) {
   printf('<tr><td>%s</td><td>', $s_key);
   if ($s_val == '') {
      $_GET['t'] = $s_key;
      $_GET['s'] = 'good';
      printf('<a href="/rating?%s">good</a>', http_build_query($_GET));
      $_GET['s'] = 'bad';
      printf('<a href="/rating?%s">bad</a>', http_build_query($_GET));
   } else {
      print $s_val;
   }
   print '</td></tr>';
}
?>
   </table>
</body>
