package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input    string
    expected []string
  }{
    {
      input: "  hello  world  ",
      expected: []string{"hello", "world"},
    },
    {
      input: "  HellO  World  ",
      expected: []string{"hello", "world"},
    },
  }

  for _, c := range cases {
    actual := cleanInput(c.input)

    if(len(actual) != len(c.expected)) {
      t.Errorf("expected %s but got %s", actual, c.expected)
      
    }
    
    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]

      if(word != expectedWord) { 
        t.Errorf("expected %s but got %s", expectedWord, word)
      }
    }
  }
}