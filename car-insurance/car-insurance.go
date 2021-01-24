package main

import (
   "fmt"
   "time"
)

const (
   agent = "QUEBEC ROMEO"
   agentPhone = "(456) 789-0123"
   company = "INDIA JULIET KILO LIMA"
   companyDept = "MIKE NOVEMBER OSCAR PAPA"
   companyPhone = "(345) 678-9012"
   driver = "ALFA B CHARLIE"
   driverAddr = "1234 DELTA ECHO FOXTROT 567"
   driverCity = "GOLF, HO 89012"
   policy = "45 SIE - 678901234"
   vehicle = "2016 TANGO UNIFORM"
   vehicleNum = "VICTO56789R012345"
)

func main() {
   from := time.Now()
   for n := 12; n > 0; n-- {
      fmt.Println(from.Format(time.RFC3339[:10]))
      from = from.AddDate(0, 1, 0)
   }
}
