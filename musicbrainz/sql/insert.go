package statement

import (
   "database/sql"
   "log"
   _ "github.com/mithrandie/csvq-driver"
)

func Insert() {
   open_o, e := sql.Open("csvq", "")
   if e != nil {
      log.Fatal(e)
   }
   var res_o sql.Result
   res_o, e = open_o.Exec("SET @@ANSI_QUOTES TO TRUE;")
   if e != nil {
      log.Fatal(res_o, e)
   }
   s := `create table "artist_t.csv" (artist_n, artist_s);
insert into artist_t values (
   (select coalesce(max(artist_n),0)+1 from artist_t), 'Goldfrapp'
);
insert into artist_t values (
   (select coalesce(max(artist_n),0)+1 from artist_t), 'Kate Bush'
);`
   res_o, e = open_o.Exec(s)
   if e != nil {
      log.Fatal(res_o, e)
   }
}
