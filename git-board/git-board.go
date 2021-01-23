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

function color_f(bool $test_b, string $mesg_s): string {
   return $test_b ? color_green($mesg_s) : color_red($mesg_s);
}

$min_n = 64;
$cha_n = $add_n = $rem_n = 0;
system('git add .');

if (is_file('config.toml')) {
   exec('git diff --cached --numstat :!docs', $tab_a);
} else {
   exec('git diff --cached --numstat', $tab_a);
}

foreach ($tab_a as $row_s) {
   $cha_n++;
   $row_a = sscanf($row_s, "%d\t%d");
   if (str_starts_with($row_s, '-')) {
      continue;
   }
   $add_n += $row_a[0];
   $rem_n += $row_a[1];
}

echo 'minimum: ', $min_n, "\n";
echo color_f($cha_n >= $min_n, 'changed files: ' . $cha_n), "\n";
echo color_f($add_n >= $min_n, 'additions: ' . $add_n), "\n";
echo color_f($rem_n >= $min_n, 'deletions: ' . $rem_n), "\n";
echo "\n";

$then_s = shell_exec('git log -1 --date=short --format=%cd');
$zone_o = new DateTimeZone('America/Chicago');
$date_o = new DateTime(timezone: $zone_o);
$now_s = $date_o->format('Y-m-d');

echo 'last commit date: ', $then_s;
echo color_f($now_s > $then_s, 'current date: ' . $now_s), "\n";
