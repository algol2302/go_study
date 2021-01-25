package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	c1 := truck{
		vehicle: vehicle{
			doors: 6,
			color: "gray",
		},
		fourWheel: false,
	}

	c2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(c1, c2)
	fmt.Println(c1.doors)
	fmt.Println(c2.doors)
}
