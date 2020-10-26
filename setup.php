<?php
declare(strict_types = 1);
error_reporting(E_ALL);

$s_to = 'C:\php\pear\sienna';

# unlink
if (is_dir($s_to)) {
   $o_iter = new FilesystemIterator($s_to);
   foreach ($o_iter as $o_info) {
      $s_path = $o_info->getPathname();
      echo 'unlink: ', $s_path, "\n";
      unlink($s_path);
   }
} else {
   mkdir($s_to);
}

# copy
$f_filter = fn ($o_info) => $o_info->getFilename() != '.git';
$o_dir = new RecursiveDirectoryIterator('.');
$o_filter = new RecursiveCallbackFilterIterator($o_dir, $f_filter);
$o_iter = new RecursiveIteratorIterator($o_filter);

foreach ($o_iter as $o_info) {
   if ($o_info->getExtension() != 'php') {
      continue;
   }
   $s_file = $o_info->getFilename();
   if ($s_file == 'setup.php') {
      continue;
   }
   $s_path = $o_info->getPathname();
   echo 'copy: ', $s_path, "\n";
   copy($s_path, $s_to . DIRECTORY_SEPARATOR . $s_file);
}
