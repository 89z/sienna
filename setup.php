<?php
declare(strict_types = 1);
error_reporting(E_ALL);

$to_s = 'C:\php\pear\sienna';

# unlink
if (is_dir($to_s)) {
   $iter_o = new FilesystemIterator($to_s);
   foreach ($iter_o as $info_o) {
      $path_s = $info_o->getPathname();
      echo 'unlink: ', $path_s, "\n";
      unlink($path_s);
   }
} else {
   mkdir($to_s);
}

# copy
$filter_f = fn ($info_o) => $info_o->getFilename() != '.git';
$dir_o = new RecursiveDirectoryIterator('.');
$filter_o = new RecursiveCallbackFilterIterator($dir_o, $filter_f);
$iter_o = new RecursiveIteratorIterator($filter_o);

foreach ($iter_o as $info_o) {
   if ($info_o->getExtension() != 'php') {
      continue;
   }
   $file_s = $info_o->getFilename();
   if ($file_s == 'setup.php') {
      continue;
   }
   $path_s = $info_o->getPathname();
   echo 'copy: ', $path_s, "\n";
   copy($path_s, $to_s . DIRECTORY_SEPARATOR . $file_s);
}
