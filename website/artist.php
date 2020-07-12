<?php
declare(strict_types = 1);
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <?= '<title>' . $s_artist . ' - Sienna</title>' ?>
</head>
<body>
   <header>
      <div>
         <a href="..">Up</a>
      </div>
   </header>
   <main>
<?php
$s_get = file_get_contents('../json/' . $s_file);
$o_get = json_decode($s_get);
foreach ($o_get->$s_artist as $s_album => $o_album) {
echo <<<eof
<div>
   <a href="/album.php?f=$s_file&a=$s_artist&r=$s_album">$s_album</a>
</div>

eof;
}
?>
   </main>
</body>
