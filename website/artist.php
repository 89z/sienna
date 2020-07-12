<?php
declare(strict_types = 1);
?>
<link rel="stylesheet" href="/sienna.css">
<main>
<?php
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
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
