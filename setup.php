<?php
declare(strict_types = 1);

$s_dir = 'C:/path/sienna';

# remove php
$a = glob($s_dir . '/*.php');
foreach ($a as $s) {
   echo 'unlink: ', $s, "\n";
   unlink($s);
}

# remove ps1
echo "unlink: sienna.ps1\n";
unlink($s_dir . '/sienna.ps1');

# add php
$a = glob('*.php');
foreach ($a as $s) {
   if ($s == 'setup.php') {
      continue;
   }
   echo 'copy: ', $s, "\n";
   copy($s, $s_dir . '/' . $s);
}

# add ps1
echo "copy: sienna.ps1\n";
copy('sienna.ps1', $s_dir . '/sienna.ps1');
