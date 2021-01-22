<?php
declare(strict_types = 1);

if ($argc != 2) {
   echo <<<eof
usage:
   go-deps.php <URL>

example:
   go-deps.php https://github.com/dinedal/textql
eof;
   exit(1);
}

$url_s = $argv[1];
$url_m = parse_url($url_s);
$mod_s = $url_m['host'] . $url_m['path'];
system('go mod init deps');
system('go get ' . $mod_s);
exec('go list -deps ' . $mod_s . '/...', $dep_a);
unlink('go.sum');
unlink('go.mod');
$dep_n = 0;

foreach ($dep_a as $dep_s) {
   if (str_contains($dep_s, '/internal/')) {
      continue;
   }
   if (! str_contains($dep_s, '.')) {
      continue;
   }
   if (str_starts_with($dep_s, $mod_s . '/')) {
      continue;
   }
   echo $dep_s, "\n";
   $dep_n++;
}

echo "\n", $dep_n, " deps\n";
