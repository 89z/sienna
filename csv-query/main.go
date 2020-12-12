package main

import (
   "database/sql"
   "log"
   _ "github.com/mithrandie/csvq-driver"
)

func Query() error {
   o, e := sql.Open("csvq", ".")
   if e != nil {
      return e
   }
   query_s := "SELECT first_name, country_code FROM users WHERE id = 12"
   var country_s, name_s string
   e = o.QueryRow(query_s).Scan(&name_s, &country_s)
   if e != nil {
      return e
   }
   println(name_s, country_s)
   return nil
}
