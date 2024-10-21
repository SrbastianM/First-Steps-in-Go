package main

// And other topic, you can add alias to the import's like: pk "learning_go/src/mypackage"
// so go on, when you call the import in your code you can access typing pk and not mypackage
import (
	"fmt"
	pk "learning_go/src/mypackage"
)

type car struct {
	brand string
	year  int
}

func main() {
	// create struct first method
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// second method
	var newCar car
	newCar.brand = "chevrolet"
	newCar.year = 2003
	fmt.Println(newCar)
	// import a public struct
	var newCarPublic = pk.CarPublic{Brand: "chevrolet", Year: 2003}
	fmt.Println(newCarPublic)

	pk.PrintMessage()

}
