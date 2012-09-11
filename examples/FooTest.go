package main

import "fmt"
import "github.com/daviejaneway/C4G/src"
import "github.com/daviejaneway/C4GUnit/src"

var Session = C4GUnit.TestSession{}

var t = C4GUnit.TestContract{
  Conditions:[]C4G.Condition{
    C4G.Condition{"Assert: @i > 0"},
    C4G.Condition{"Assert: @i * @i == 25"}}}

func foo(i int) int {
  return i * i
}

func testFoo() {
  i := foo(10)
  
  t.Assert(0, i > 0)
}

func testFoo2() {
  i := foo(6)
  
  t.Assert(1, i == 25)
}

func main() {
  testFoo()
  testFoo2()
  
  fmt.Println(&Session)
}