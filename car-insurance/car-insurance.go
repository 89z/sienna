package main

const (
   address1 = "1234 DELTA ECHO FOXTROT 567"
   address2 = "GOLF, HO 89012"
   agent_name = "QUEBEC ROMEO"
   agent_phone = "(456) 789-0123"
   company1 = "INDIA JULIET KILO LIMA"
   company2 = "MIKE NOVEMBER OSCAR PAPA"
   company_phone = "(345) 678-9012"
   driver_name = "ALFA B CHARLIE"
   make_model = "2016 TANGO UNIFORM"
   policy_num = "45 SIE - 678901234"
   vehicle_num = "VICTO56789R012345"
)

<!doctype html>
<html lang="en">
<head>
   <link rel="stylesheet" href="cove.css">
   <meta charset="utf-8">
   <title>Car insurance</title>
</head>
<body>
$date_o = new DateTime;
$add_o = new DateInterval('P1M');
for ($n = 12; $n > 0; $n--) {
   $from_s = $date_o->format('Y-m-d');
   $date_o->add($add_o);
   $to_s = $date_o->format('Y-m-d');
}
</body>
</html>
