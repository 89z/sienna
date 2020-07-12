<?php
declare(strict_types = 1);
?>
<style>
div {
   background: burlywood;
   margin: 1em;
   padding: 1em;
}
main {
   display: flex;
   flex-wrap: wrap;
}
</style>
<main>
<?php
$s_file = $_GET['f'];
$s_get = file_get_contents('D:/Git/umber/data/' . $s_file);
$o_get = json_decode($s_get);

foreach ($o_get as $s_artist => $o_artist) {
   echo <<<eof
<div>$s_artist</div>

eof;
}
?>
</main>


