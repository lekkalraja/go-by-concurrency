package main

func main() {
	var ch chan string

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send (nil chan)]:
	*/
	// ch <- "data"

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive (nil chan)]:
	*/
	//<-ch

	/*
		panic: close of nil channel

		goroutine 1 [running]:
	*/
	//close(ch)

	ch = make(chan string)
	close(ch)
	/*
		panic: close of closed channel

		goroutine 1 [running]:
	*/
	// close(ch)

}
