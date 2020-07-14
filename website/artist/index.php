<?php
declare(strict_types = 1);
require '../sienna.php';
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
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

foreach ($o_get->$s_artist as $s_album => $o_album) {
   echo '<tr>';
   if ($s_album[0] == '@') {
      echo <<<eof
<td>$s_album</td>
<td>$o_album</td>
eof;
   } else {
      $s_date = $o_album->{'@date'};
      $s_class = $m_local[$s_album];
      $_GET['r'] = $s_album;
      $s_q = http_build_query($_GET);
      echo <<<eof
<td>$s_date</td>
<td class="$s_class">
   <a href="/release?$s_q">$s_album</a>
</td>
eof;
   }
   echo '</tr>';
}
?>
   </table>
   <footer>
<?php
echo <<<eof
<a href="/add-release?f=$s_file&a=$s_artist">Add release</a>
eof;
?>
   </footer
</body>
