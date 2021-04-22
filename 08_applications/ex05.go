package main

import (
	"fmt"
	"sort"
)

type user5 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// By defines a sorting method
type By func(u1, u2 *user5) bool

type userSorter struct {
	users []user5
	by    func(u1, u2 *user5) bool
}

// Sort defines an user sorter
func (by By) Sort(users []user5) {
	us := &userSorter{
		users: users,
		by:    by,
	}
	sort.Sort(us)
}

func (u *userSorter) Len() int {
	return len(u.users)
}

func (u *userSorter) Swap(i, j int) {
	u.users[i], u.users[j] = u.users[j], u.users[i]
}

func (u *userSorter) Less(i, j int) bool {
	return u.by(&u.users[i], &u.users[j])
}

func main() {
	u1 := user5{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user5{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user5{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user5{u1, u2, u3}

	// fmt.Println(users)

	age := func(u1, u2 *user5) bool {
		return u1.Age < u2.Age
	}

	last := func(u1, u2 *user5) bool {
		return u1.Last < u2.Last
	}

	By(age).Sort(users)
	//fmt.Println("By age:", users)

	By(last).Sort(users)
	//fmt.Println("By last name:", users)

	for _, u := range users {
		fmt.Printf("%v %v\n", u.First, u.Last)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}

}
