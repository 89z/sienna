<?php
declare(strict_types = 1);

system('git commit --verbose');

if (is_file('config.toml')) {
   $dir_o = new RecursiveDirectoryIterator('docs');
   $iter_o = new RecursiveIteratorIterator($dir_o);

   echo "UNLINK\n";
   foreach ($iter_o as $info_o) {
      if ($info_o->isFile()) {
         unlink($info_o->getPathname());
      }
   }

   echo "HUGO\n";
   system('hugo');
   echo "GIT ADD\n";
   system('git add .');
   echo "GIT COMMIT\n";
   system('git commit --amend');
}

echo "git push\n";
