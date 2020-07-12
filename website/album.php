<?php
declare(strict_types = 1);
?>
<link rel="stylesheet" href="/sienna.css">
<main>
<?php
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
$s_rel = $_GET['r'];
$s_get = file_get_contents('../json/' . $s_file);
$o_get = json_decode($s_get);
foreach ($o_get->$s_artist->$s_rel as $s_track => $s_rate) {
echo <<<eof
<div>$s_track</div>

eof;
}
?>
</main>
