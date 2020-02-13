package main

import (
	"fmt"
	"sort"
)

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user2{u1, u2, u3}

	fmt.Println("Unsorted:", users)

	sort.Sort(ByAge(users))
	fmt.Println("-----")
	fmt.Println("Sorted By Age:")
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("saying:", saying)
		}
		fmt.Println("------")
	}

	sort.Sort(ByLast(users))
	fmt.Println("-----")
	fmt.Println("Sorted By Last:")
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("saying:", saying)
		}
		fmt.Println("------")
	}

	for _, user := range users {
		fmt.Println("--------")
		fmt.Println("Sorted sayings for user", user.First, user.Last)
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Println(saying)
		}
		fmt.Println("--------")
	}

}

type ByAge []user2

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user2

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }
