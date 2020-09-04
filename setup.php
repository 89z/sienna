<?php
declare(strict_types = 1);
error_reporting(E_ALL);

$s_to = 'C:/php/pear/sienna';

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

# copy bin
$o_iter = new FilesystemIterator('bin');

foreach ($o_iter as $o_info) {
   $s_file = $o_info->getFilename();
   if ($s_file == 'readme.md') {
      continue;
   }
   $s_path = $o_info->getPathname();
   echo 'copy: ', $s_path, "\n";
   copy($s_path, $s_to . '/' . $s_file);
}

# copy include
$o_iter = new FilesystemIterator('include');

foreach ($o_iter as $o_info) {
   $s_file = $o_info->getFilename();
   if ($s_file == 'readme.md') {
      continue;
   }
   $s_path = $o_info->getPathname();
   echo 'copy: ', $s_path, "\n";
   copy($s_path, $s_to . '/' . $s_file);
}
