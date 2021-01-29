package main

var tomlSep = []byte{'+', '+', '+', '\n'}

type frontMatter struct{
   Build struct{
      List string
   } `toml:"_build"`
   Example []string
}
