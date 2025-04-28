package main 

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input     string
    expected  []string
  }{
    {
      input:    "  hello word    ",
      expected: []string{"hello", "word"},
    },
  }

  for _, c := range cases {
    actual := cleanInput(c.input)

    if len(actual) != len(c.expected) {
       t.Errorf("\nLength\nExpected: %d\nActual: %d\n", len(c.expected), len(actual))
    }
    
    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]
      if word != expectedWord {
        t.Errorf("Word: %s is not equal to expectedWord: %s", word, expectedWord)
      }
    }
  }
}
