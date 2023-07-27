package main

import "fmt"

func fibonacci(n int)int{
	if n<=1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2);
}

func main()  {
	num := 8
	fmt.Println("O" ,num, "numero na sequencia de fibonacci é: ", fibonacci(num))
}