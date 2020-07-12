<?php
declare(strict_types = 1);
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
$s_rel = $_GET['r'];
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <?= '<title>' . $s_rel . ' - Sienna</title>' ?>
</head>
<body>
   <header>
<?php
echo <<<eof
<a href="/artist.php?f=$s_file&a=$s_artist">Up</a>
<h1>$s_artist - $s_rel</h1>
eof;
?>
   </header>
   <table>
<?php
$s_get = file_get_contents('../json/' . $s_file);
$o_get = json_decode($s_get);
foreach ($o_get->$s_artist->$s_rel as $s_key => $s_val) {
   echo <<<eof
<tr>
   <td>$s_key</td>
   <td>$s_val</td>
</tr>
eof;
}
?>
   </table>
</body>
