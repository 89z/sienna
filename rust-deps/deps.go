package main

var (
   dep int
   prev string
)

type m map[string]interface{}

type cargoLock struct{
   Package []struct{
      Name string
   }
}
