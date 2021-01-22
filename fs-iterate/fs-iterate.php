<?php
declare(strict_types = 1);

function color_cyan(string $s): string {
   return "\e[1;36m" . $s . "\e[m";
}

function color_green(string $s): string {
   return "\e[1;32m" . $s . "\e[m";
}

function color_red(string $s): string {
   return "\e[1;31m" . $s . "\e[m";
}

if ($argc < 3) {
   echo "fs-iterate.php <path> <command-line>\n";
   exit(1);
}

# part 1
$cmd_a = array_slice($argv, 2);
$dir_s = $argv[1];
# part 2
$desc_a = [];
$pipe_a = [];
$dir_o = new FilesystemIterator($dir_s);
$cmd_s = implode(' ', $cmd_a);

foreach ($dir_o as $ent_o) {
   $path_s = $ent_o->getPathname();
   chdir($path_s);
   echo color_cyan($path_s), "\n";
   # system loses the color
   $r = proc_open($cmd_s, $desc_a, $pipe_a);
   proc_close($r);
}
