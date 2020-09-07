<?php
declare(strict_types = 1);
error_reporting(E_ALL);

if ($argc != 2) {
   echo "pretty.php <file>";
   exit(1);
}

$s_file = $argv[1];
$s_get = file_get_contents($s_file);
$o_json = json_decode($s_get);
$n_err = json_last_error();

if ($n_err != JSON_ERROR_NONE) {
   echo "invalid JSON\n";
   exit(1);
}

$n = JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE;
$s_json = json_encode($o_json, $n);
file_put_contents($s_file, $s_json . "\n");
