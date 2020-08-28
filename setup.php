<?php
declare(strict_types = 1);

$s_dir = 'C:/php/pear/sienna';

# unlink
if (is_dir($s_dir)) {
   $a = glob($s_dir . '/*');
   foreach ($a as $s) {
      echo 'unlink: ', $s, "\n";
      unlink($s);
   }
} else {
   mkdir($s_dir);
}

# copy bin
$a = glob('bin/*');

foreach ($a as $s) {
   $s_base = basename($s);
   if ($s_base == 'readme.md') {
      continue;
   }
   echo 'copy: ', $s, "\n";
   copy($s, $s_dir . '/' . $s_base);
}

# copy include
$a = glob('include/*.php');

foreach ($a as $s) {
   echo 'copy: ', $s, "\n";
   $s_base = basename($s);
   copy($s, $s_dir . '/' . $s_base);
}
