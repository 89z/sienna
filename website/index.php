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
   echo <<<eof
<div>
   <a href="/json.php?f=$s_ent">$s_ent</a>
</div>

eof;
}
?>
   </main>
</body>
