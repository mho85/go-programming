package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

func main() {
	myTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(myTruck)
	fmt.Println(mySedan)
	fmt.Println(myTruck.fourWheel)
	fmt.Println(mySedan.luxury)

}
