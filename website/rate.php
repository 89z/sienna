<?php
declare(strict_types = 1);
$s_file = $_GET['f'];
$s_artist = $_GET['a'];
$s_rel = $_GET['r'];
$s_track = $_GET['t'];
$s_rate = $_GET['s'];
$s_get = file_get_contents('../json/' . $s_file . '.json');
$o_get = json_decode($s_get);
$o_get->$s_artist->$s_rel->$s_track = $s_rate;
$s_json = json_encode($o_get, JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);
file_put_contents($s_file . '.json', $s_json);
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <title>Sienna</title>
</head>
<body>
   <h1>Complete</h1>
</body>
