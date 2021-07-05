package msi

import (
   "database/sql"
   "fmt"
   "testing"
)

const kit = "Windows App Certification Kit SupportedApiList ARM-arm_en-us.msi"

func TestOne(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT Cabinet FROM Media")
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   for rows.Next() {
      var cabinet string
      rows.Scan(&cabinet)
      fmt.Printf("%q\n", cabinet)
   }
}

func TestTwo(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT FileName FROM File")
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   for rows.Next() {
      var file string
      rows.Scan(&file)
      fmt.Printf("%q\n", file)
   }
}

func TestThree(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT DiskId,Cabinet FROM Media")
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   for rows.Next() {
      var (
         diskId int
         cabinet string
      )
      rows.Scan(&diskId, &cabinet)
      fmt.Printf("%v %q\n", diskId, cabinet)
   }
}

func TestFour(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT DiskId,Cabinet FROM Media")
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   col, err := rows.Columns()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", col)
}

func TestFive(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT ? FROM Media", "Cabinet")
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   for rows.Next() {
      var cabinet string
      rows.Scan(&cabinet)
      fmt.Printf("%q\n", cabinet)
   }
}

func TestSix(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query(
      `SELECT DiskId FROM Media WHERE Cabinet = ?`,
      "3cf96a08c3b29e9dcf5946d28affb747.cab",
   )
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   rows.Next()
   var diskId int
   rows.Scan(&diskId)
   println(diskId)
}

func TestSeven(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT Cabinet FROM Media WHERE DiskId = 2")
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   rows.Next()
   var cabinet string
   rows.Scan(&cabinet)
   println(cabinet)
}

func TestEight(t *testing.T) {
   db, err := sql.Open("msi", kit)
   if err != nil {
      t.Fatal(err)
   }
   defer db.Close()
   rows, err := db.Query("SELECT Cabinet FROM Media WHERE DiskId = ?", 2)
   if err != nil {
      t.Fatal(err)
   }
   defer rows.Close()
   rows.Next()
   var cabinet string
   rows.Scan(&cabinet)
   println(cabinet)
}
