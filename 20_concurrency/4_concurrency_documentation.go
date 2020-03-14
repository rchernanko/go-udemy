package main

/*
	https://golang.org/doc/effective_go.html#concurrency

	Key highlights:

	Concurrency:

	- "Concurrent programming in many environments is made difficult by the subtleties required to implement correct access to shared variables."
	- "Go encourages a different approach in which shared values are passed around on channels and, in fact, never
		actively shared by separate threads of execution. Only one goroutine has access to the value at any given time.
			Data races cannot occur, by design. To encourage this way of thinking we have reduced it to a slogan:
				Don't communicate by sharing memory; instead, share memory by communicating.

				Great video from Rob Pike on the origin of the above phrase - https://www.youtube.com/watch?v=f6kdp27TYZs&feature=youtu.be

		Basically "Don't communicate by sharing memory; instead, share memory by communicating" = you don't have some
		blob of memory and then put locks and mutexes around it to protect it from parallel access. Instead you use the
		channel to pass the data back and forth between the go routines.

	- i.e. use channels to manage the data between several go routines

	Go Routines:

	- "A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space"
	- "Prefix a function or method call with the go keyword to run the call in a new goroutine.
		When the call completes, the goroutine exits, silently. (The effect is similar to the Unix shell's & notation
			for running a command in the background.)"

	https://golang.org/ref/spec#Go_statements

	- "A "go" statement starts the execution of a function call as an independent concurrent thread of control, or
		goroutine, within the same address space."
	- "The function value and parameters are evaluated as usual in the calling goroutine, but unlike with a regular call,
		program execution does not wait for the invoked function to complete. Instead, the function begins executing
			independently in a new goroutine. When the function terminates, its goroutine also terminates. If the
				function has any return values, they are discarded when the function completes."
	- re: above, in general, you don't really launch something into a go routine where you're looking to have that function return something back
		Instead you could write the return value to a channel, and then access the return value from there e.g. https://play.golang.org/p/lBKFKCwrue


*/
