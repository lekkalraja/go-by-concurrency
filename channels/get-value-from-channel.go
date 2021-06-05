package main

import "fmt"

func main() {

	ch := make(chan int)

	go func(num1, num2 int) {
		defer close(ch)
		sum := num1 + num2
		ch <- sum
		fmt.Printf("[Go-Routine] Calculated sum for %d, %d : %d \n", num1, num2, sum)
		// return sum // ERROR : Should not return from go routine
	}(2, 4)

	value, ok := <-ch // It will receive whenever sender send's the data

	fmt.Printf("[Main-Routine] Received sum : %d and status : %v\n", value, ok)

	value, ok = <-ch // It will receive once the close operation done.

	fmt.Printf("[Main-Routine] Received sum : %d and status : %v\n", value, ok)
}

/*
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run channels/get-value-from-channel.go
	[Go-Routine] Calculated sum for 2, 4 : 6
	[Main-Routine] Received sum : 6 and status : true
	[Main-Routine] Received sum : 0 and status : false
*/
