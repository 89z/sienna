<?php
declare(strict_types = 1);
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <title>Sienna</title>
</head>
<body>
   <main>
<?php
$a_scan = scandir('../json');
foreach ($a_scan as $s_ent) {
   if ($s_ent == '.') {
      continue;
   }
   if ($s_ent == '..') {
      continue;
   }
   $s_get = file_get_contents('../json/' . $s_ent);
   $o_get = json_decode($s_get);
   foreach ($o_get as $s_artist => $o_artist) {
echo <<<eof
<div>
   <a href="/artist.php?f=$s_ent&a=$s_artist">$s_artist</a>
</div>

eof;
   }
}
?>
   </main>
</body>
