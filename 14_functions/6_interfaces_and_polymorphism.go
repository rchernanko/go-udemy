package main

import "fmt"

/**

When we define the below 'human' interface, we are saying that any 'type' that has the method 'speakPlease' is also
of type 'human' i.e. they implement the 'human' interface

'If you have my method, then you're also of my type'. i.e. 'if you have the speakPlease method, you're also of type human'

TODO - He recommends to read this - https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

*/

type human interface {
	speakPlease()
}

type person1 struct {
	first string
	last  string
}

func (p person1) speakPlease() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type secretAgent1 struct {
	person1
	licenseToKill bool
}

func (sa secretAgent1) speakPlease() {
	fmt.Println("I am", sa.first, sa.last, " - the secretAgent speak")
}

//HAVE ADDED ASSERTION BELOW - e.g. I am asserting that h is of type 'person1'
func bar5(h human) {

	switch h.(type) {
	case person1:
		fmt.Println("I was passed into bar5", h.(person1).first) //So you can do this when you have an interface but you then need to get back to the concrete 'type'
	case secretAgent1:
		fmt.Println("I was passed into bar5", h.(secretAgent1).first)
	}
}

// IMPORTANT - A VALUE CAN BE OF MORE THAN ONE TYPE. In the below, the value sa1 is of type 'secretAgent' AND 'human'

func main() {
	sa1 := secretAgent1{
		person1: person1{
			first: "miss",
			last:  "moneypenny",
		},
		licenseToKill: false,
	}

	sa1.speakPlease()
	bar5(sa1)

	p1 := person1{
		first: "dr",
		last:  "yes",
	}

	p1.speakPlease()
	bar5(p1)
}
