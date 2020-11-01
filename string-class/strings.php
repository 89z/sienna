<?php
declare(strict_types = 1);
error_reporting(E_ALL);

function str_contains($s, $s2) {
   return strpos($s, $s2) !== false;
}
