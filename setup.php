<?php
declare(strict_types = 1);

$s_dir = 'C:/path/cove';

if (is_dir($s_dir)) {
   $o_dir = new RecursiveDirectoryIterator($s_dir);
   $o_iter = new RecursiveIteratorIterator($o_dir);
   foreach ($o_iter as $o_path) {
      if (! $o_path->isDir()) {
         $s_rm = $o_path->getPathname();
         echo 'unlink: ', $s_rm, "\n";
         unlink($s_rm);
      }
   }
} else {
   mkdir($s_dir);
}

$o_dir = new RecursiveDirectoryIterator('.');
$o_iter = new RecursiveIteratorIterator($o_dir);

foreach ($o_iter as $o_path) {
   $s_base = $o_path->getBasename();
   if ($s_base == 'setup.php') {
      continue;
   }
   $n_mat = preg_match('/\.php$/', $s_base);
   if ($n_mat === 0) {
      continue;
   }
   $s_path = $o_path->getPathname();
   echo 'copy: ', $s_path, "\n";
   copy($s_path, $s_dir . '/' . $s_base);
}
