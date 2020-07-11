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
$a_scan = scandir('D:/Git/umber/data');
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
