package main

import (
	"computorv1/src"
	"fmt"
	"os"
)

// aX^2 + bX + c
func main() {
  if len(os.Args) != 2 {
    fmt.Println("Wrong number of arguments")
    return 
  }

  str := os.Args[1]

  if str == "" {
    fmt.Println("Empty string, need one expresion like 'a * X^2 + b * X^1 + c * X^0 = ...'")
    return
  }

  exp, err := src.ParseExpression(str)
  
  if err != nil {
    fmt.Println(err)
    return
  }

  exp.PrintReducedForm()
  exp.PrintPolDegree()
  exp.PrintSolution()
}

