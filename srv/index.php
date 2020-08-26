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
$o_dir = new FilesystemIterator('data');

foreach ($o_dir as $o_ent) {
   $s_ent = $o_ent->getFilename();
   $s_get = file_get_contents('data/' . $s_ent);
   $o_get = json_decode($s_get);
   $m_q['f'] = pathinfo($s_ent, PATHINFO_FILENAME);
   foreach ($o_get as $s_artist => $o_artist) {
      $m_q['a'] = $s_artist;
      $s_q = http_build_query($m_q);
      printf('<div><a href="/artist?%s">%s</a></div>', $s_q, $s_artist);
   }
}
?>
   </main>
</body>
