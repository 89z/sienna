<?php
declare(strict_types = 1);

# part 1
$s_artist = $_GET['a'];
$s_file = $_GET['f'];
$s_rel = $_GET['r'];
$s_rate = $_GET['s'];
$s_track = $_GET['t'];
# part 2
$s_get = file_get_contents('../json/' . $s_file . '.json');
# part 3
$n_json = JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES;
$o_get = json_decode($s_get);
$o_get->$s_artist->$s_rel->$s_track = $s_rate;
$s_json = json_encode($o_get, $n_json);
file_put_contents('../json/' . $s_file . '.json', $s_json . "\n");
?>
<head>
   <link rel="icon" href="/sienna.png">
   <link rel="stylesheet" href="/sienna.css">
   <title>Sienna</title>
</head>
<body>
   <header>
<?php
echo <<<eof
<a href="/release?f=$s_file&a=$s_artist&f=$s_rel">Back</a>
eof;
?>
   <h1>Complete</h1>
   </header>
</body>
