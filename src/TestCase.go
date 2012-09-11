package C4GUnit

import "github.com/daviejaneway/C4G/src"

type TestSession struct {
  passes int
  failures int
}

var Session = TestSession{}

type TestContract struct {
  Conditions []C4G.Condition
}

func(tc TestContract) Assert(index int, b bool) {
  defer func() {
      if r := recover(); r != nil {
        Session.failures++
      } else {
        Session.passes++
      }
  }()
  
  tc.Conditions[index].Execute(b)
}