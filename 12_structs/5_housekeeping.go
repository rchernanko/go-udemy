package main

/*

Some nice to know stuff:

1)

It's all about type
Is go an object oriented language? Go has OOP aspects. But it’s all about TYPE.
We create TYPES in Go; user-defined TYPES. We can then have VALUES of that type. We can assign VALUES of a user-defined
TYPE to VARIABLES.

2)

Go is Object Oriented
	a) Encapsulation
		- state ("fields")
		- behavior ("methods")
		- exported & unexported; viewable & not viewable
	b) Reusability
		- inheritance ("embedded types")
	c) Polymorphism
		- interfaces
	d) Overriding
		-"promotion"

3)

Traditional OOP
	a) Classes
		- data structure describing a type of object
		- you can then create "instances"/"objects" from the class / blueprint
		- classes hold both:
			- state / data / fields
			- behavior / methods
		- public / private
	b) Inheritance

4)

In Go:
	- you don't create classes, you create a TYPE
	- you don't instantiate, you create a VALUE of a TYPE

5)

User defined types
 	- We can declare a new type
	- e.g. 'type foo int'
		- the underlying type of foo is int
		- int conversion
			- int(myAge)
			- converting type foo to type int

	- IT IS A BAD PRACTICE TO ALIAS TYPES
		- one exception:
			- if you need to attach methods to a type
			- see the time package for an example of this godoc.org/time
				- type Duration int64
				- Duration has methods attached to it


TODO up to 11:20

*/
