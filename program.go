package main

import "os"
import "fmt"
import "io/ioutil"
import "github.com/moovweb/gokogiri/xml"

type Test struct {
}


func main() {
  args := os.Args[1:]

  dat, err := ioutil.ReadFile(args[0])
  if err != nil {
    fmt.Println(err)
  }

  doc, err := xml.Parse(dat, xml.DefaultEncodingBytes, nil, xml.StrictParseOption, xml.DefaultEncodingBytes)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  } else {
    doc.Free()
  }

  os.Exit(0)
}
