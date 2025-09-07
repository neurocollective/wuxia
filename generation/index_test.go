package generation

import (
  "testing"
)

func TestSQLArgSequence(t *testing.T) {

  schema, err := ReadDump()

  if string != expected {
    t.Log("got:\n", string)
    t.Log("expected:\n", expected)
    t.Fatal("incorrect result")
  }

}