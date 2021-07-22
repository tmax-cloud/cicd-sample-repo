package main

import ( "fmt"
         "rsc.io/quote"
		 "test/pkg/calc"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(calc.Sum(1,2))
}
