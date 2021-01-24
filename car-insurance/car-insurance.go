package main

import (
   "fmt"
   "html/template"
   "log"
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

const tmpl = `<h1>{{ .agent }}</h1>
<h2>{{ .agentPhone }}</h2>
`

func main() {
   from := time.Now()
   insure["months"] = []string{}
   for n := 12; n > 0; n-- {
      month := from.String()[:10]
      insure["months"] = append(insure["months"].([]string), month)
      from = from.AddDate(0, 1, 0)
   }
   fmt.Println(insure)
   t, e := template.New("index").Parse(tmpl)
   if e != nil {
      log.Fatal(e)
   }
   e = t.Execute(os.Stdout, insure)
   if e != nil {
      log.Fatal(e)
   }
}
