// Database SQL driver
package msi
//go:generate mkwinsyscall -output zmsi.go msi.go
//sys InstallProduct(packagePath string, command string) (e error) = msi.MsiInstallProductW
//sys closeHandle(any Conn) (e error) = msi.MsiCloseHandle
//sys createRecord(params int) (n int) = msi.MsiCreateRecord
//sys databaseOpenView(database Conn, query string, view *Stmt) (e error) = msi.MsiDatabaseOpenViewW
//sys openDatabase(dbPath string, persist int, database *Conn) (e error) = msi.MsiOpenDatabaseW
//sys recordGetFieldCount(record int) (n int) = msi.MsiRecordGetFieldCount
//sys recordGetString(record int, field int, buf *uint16, bufSize *int) (e error) = msi.MsiRecordGetStringW
//sys recordSetInteger(record int, field int, value int64) (e error) = msi.MsiRecordSetInteger
//sys recordSetString(record int, field int, value string) (e error) = msi.MsiRecordSetStringW
//sys viewClose(view Stmt) (e error) = msi.MsiViewClose
//sys viewExecute(view Stmt, record int) (e error) = msi.MsiViewExecute
//sys viewFetch(view Rows, record *int) (e error) = msi.MsiViewFetch
//sys viewGetColumnInfo(view Rows, columnInfo int, record *int) (e error) = msi.MsiViewGetColumnInfo
import (
   "database/sql"
   "database/sql/driver"
   "golang.org/x/sys/windows"
)

const (
   msicolinfo_names = 0
   msidbopen_readonly = 0
)

func init() {
   var dr Driver
   sql.Register("msi", dr)
}

type Driver struct{}

func (Driver) Open(name string) (driver.Conn, error) {
   var co Conn
   err := openDatabase(name, msidbopen_readonly, &co)
   if err != nil {
      return nil, err
   }
   return co, nil
}

type Conn int

func (Conn) Begin() (driver.Tx, error) {
   panic("Conn Begin")
}

func (c Conn) Close() error {
   return closeHandle(c)
}

func (c Conn) Prepare(query string) (driver.Stmt, error) {
   var st Stmt
   err := databaseOpenView(c, query, &st)
   if err != nil {
      return nil, err
   }
   return st, nil
}

type Rows int

func (Rows) Close() error {
   return nil
}

func (r Rows) Columns() []string {
   var record int
   err := viewGetColumnInfo(r, msicolinfo_names, &record)
   if err != nil {
      return nil
   }
   cols := make([]string, recordGetFieldCount(record))
   for n := range cols {
      var bufSize int
      err := recordGetString(record, n + 1, nil, &bufSize)
      if err != nil {
         return nil
      }
      bufSize++
      buf := make([]uint16, bufSize)
      if err := recordGetString(record, n + 1, &buf[0], &bufSize); err != nil {
         return nil
      }
      cols[n] = windows.UTF16ToString(buf)
   }
   return cols
}

func (r Rows) Next(dest []driver.Value) error {
   var record int
   err := viewFetch(r, &record)
   if err != nil {
      return err
   }
   for n := range dest {
      var bufSize int
      err := recordGetString(record, n+1, nil, &bufSize)
      if err != nil {
         return err
      }
      bufSize++
      buf := make([]uint16, bufSize)
      if err := recordGetString(record, n+1, &buf[0], &bufSize); err != nil {
         return err
      }
      dest[n] = driver.Value(windows.UTF16ToString(buf))
   }
   return nil
}

type Stmt int

func (s Stmt) Close() error {
   return viewClose(s)
}

func (Stmt) Exec(args []driver.Value) (driver.Result, error) {
   panic("Stmt Exec")
}

func (Stmt) NumInput() int {
   return -1
}

func (s Stmt) Query(args []driver.Value) (driver.Rows, error) {
   record := createRecord(len(args))
   for ind, val := range args {
      switch v := val.(type) {
      case string:
         err := recordSetString(record, ind + 1, v)
         if err != nil {
            return nil, err
         }
      case int64:
         err := recordSetInteger(record, ind + 1, v)
         if err != nil {
            return nil, err
         }
      }
   }
   err := viewExecute(s, record)
   if err != nil {
      return nil, err
   }
   return Rows(s), nil
}
