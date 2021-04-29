/*
.import --csv power.csv power
*/
select
   ("Price/kWh 1000" + "Price/kWh 500" * 11) as price,
   RepCompany,
   "Price/kWh 500",
   "Price/kWh 1000"
from power
where Rating > 1
and "New Customer" = 'False'
order by price;
