<?php
declare(strict_types = 1);

$s = 'go-deps.php';
echo $s, "\n";
copy($s, 'C:\php\pear' . DIRECTORY_SEPARATOR . $s);
