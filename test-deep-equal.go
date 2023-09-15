package main

import (
    "fmt"
    "reflect"
)

func changed() {
  m1 := map[string]string{}
  m2 := map[string]string{}

  m1["USERNAME"] = "Son"
  m2["USERNAME"] = "Son"

  return !reflect.DeepEqual(m1, m2)
}

func main() {

}
