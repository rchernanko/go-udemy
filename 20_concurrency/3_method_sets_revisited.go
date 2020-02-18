package main

/*

- Covered in a previous tutorial that I've not yet taken notes on...
- There are 2 types of receivers:

Receiver:		Values:

  (t T) 	    T and *T
  (t *T)		*T

- However, in the previous tutorial, we were using WaitGroup and we were actually going against the rules above...weird
- We had a variable of type sync.WaitGroup but were then calling methods that had a pointer receiver i.e wg.Add() wg.Done()
- Let's try to understand this a little more (even Todd doesn't fully get it) by looking at the documentation for method sets
	- https://golang.org/ref/spec#Method_sets
		- mmm this is not that clear to me still
- Dug around a little, looking at the Q&A section on this particular video, some people have put these (which makes more sense!):
	- https://tour.golang.org/methods/6
 	- https://tour.golang.org/methods/7
	- https://stackoverflow.com/questions/42477951/what-is-the-method-set-of-sync-waitgroup
*/
