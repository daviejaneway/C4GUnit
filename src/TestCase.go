package C4GUnit

import "fmt"
import "github.com/daviejaneway/C4G/src"

type TestSession struct {
  Passes int
  Failures int
}

func (ts TestSession) String() string {
  return fmt.Sprintf("ASSERTIONS: %d, FAILURES: %d", ts.Passes, ts.Failures)
}

var Session = TestSession{}

type TestContract struct {
  Conditions []C4G.Condition
}

func(tc TestContract) Assert(index int, b bool) {
  defer func() {
      if r := recover(); r != nil {
        Session.Failures++
      } else {
        Session.Passes++
      }
  }()
  
  tc.Conditions[index].Execute(b)
}