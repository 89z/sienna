<?php
declare(strict_types = 1);

$s_pear = 'C:/php/pear';

if (! is_dir($s_pear . '/sienna')) {
   mkdir($s_pear);
   mkdir($s_pear . '/sienna');
}

$a_php = glob('*.php');

foreach ($a_php as $s_php) {
   if ($s_php == 'setup.php') {
      continue;
   }
   copy($s_php, $s_pear . '/sienna/' . $s_php);
}
