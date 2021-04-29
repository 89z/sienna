package main

import (
   "html/template"
   "os"
   "time"
)

var insure = map[string]interface{}{
   "agent": "QUEBEC ROMEO",
   "agentPhone": "(456) 789-0123",
   "company": "INDIA JULIET KILO LIMA",
   "companyDept": "MIKE NOVEMBER OSCAR PAPA",
   "companyPhone": "(345) 678-9012",
   "driver": "ALFA B CHARLIE",
   "driverAddr": "1234 DELTA ECHO FOXTROT 567",
   "driverCity": "GOLF, HO 89012",
   "policy": "45 SIE - 678901234",
   "vehicle": "2016 TANGO UNIFORM",
   "vehicleNum": "VICTO56789R012345",
}

func main() {
   from := time.Now()
   var months [12]struct { From, To string }
   for n := range months {
      months[n].From = from.String()[:10]
      from = from.AddDate(0, 1, 0)
      months[n].To = from.String()[:10]
   }
   insure["months"] = months
   in, e := template.ParseFiles("in.html")
   if e != nil {
      panic(e)
   }
   out, e := os.Create("out.html")
   if e != nil {
      panic(e)
   }
   e = in.Execute(out, insure)
   if e != nil {
      panic(e)
   }
}
