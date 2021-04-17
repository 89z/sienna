package youtube

import (
   "strings"
   "testing"
)

func TestWorker(t *testing.T) {
   good := `<HTML>
<HEAD>
<title> Hivelocity Speedtest Server</title>
</HEAD>
Speedtest files:
<br>
10485760    : <A HREF="../10Mio.dat"> 10Mio.dat</A>
<br>
102400000   : <A HREF="../100mb.file"> 100mb.file</A>
<br>
10737418240 : <A HREF="../10Gio.dat"> 10Gio.dat</A>

</HTML>
`
   addr := "http://speedtest.lax.hivelocity.net"
   _, err := workerGet(addr, -1)
   if err == nil {
      t.Error()
   }
   _, err = workerGet(addr, 0)
   if err == nil {
      t.Error()
   }
   src, err := workerGet(addr, 2)
   if err != nil {
      t.Error(err)
   }
   dst := new(strings.Builder)
   for _, each := range src {
      each.WriteTo(dst)
   }
   if dst.String() != good {
      t.Error(dst.String())
   }
}
