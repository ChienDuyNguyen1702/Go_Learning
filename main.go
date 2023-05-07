package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/ChienDuyNguyen1702/MyFirstGoLang/helper"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

//
//interfaces
//

type Animal interface {
	Say() string
	NumberOfLegs() int
}
type Dog struct {
	Name  string
	Breed string // giống loài
}
type Cat struct {
	Color    string
	EyeColor string
}

// Naming convention:
func main() {
	fmt.Println("He no guon")

	var whatToSay string
	var i int

	whatToSay = "goodbye im dead"
	fmt.Println(whatToSay)

	i = 2
	fmt.Println("i equals ", i)
	// fmt.Println(saySomething())
	whatWasSay := saySomething()
	fmt.Println(whatWasSay)

	//
	// pointers
	//
	// var color string
	color := "green"
	fmt.Println(color)
	changeThing(&color)
	fmt.Println(color)

	//
	// Maps
	//
	myMap := make(map[string]User) //khai bao maps

	myMap["me"] = User{ // gán maps cách 1
		FirstName: "Chien",
		LastName:  "Nguyen",
	}
	KhanhHuyen := User{"Huyen", "btk", "mei", 19} // gán maps cách 2
	myMap["kh"] = KhanhHuyen
	fmt.Println(myMap["me"].FirstName)
	fmt.Println(myMap["kh"].FirstName)
	fmt.Println(myMap["kh"])

	//
	//slices is similar to arrays
	//
	var animal []string
	animal = append(animal, "fish")
	animal = append(animal, "cat")
	animal = append(animal, "dog")
	animal = append(animal, "bird")
	fmt.Println(animal)

	var number []int
	number = append(number, 3)
	number = append(number, 1)
	number = append(number, 2)
	number = append(number, -2)
	fmt.Println(number)
	sort.Ints(number)
	fmt.Println(number)

	number2 := []int{4, 6, 26, 78, 853, 365, 7, 2, 64, 0}

	fmt.Println(number2)
	sort.Ints(number2)
	fmt.Println(number2)

	fmt.Println(number2[4:5]) //slice[start_index:end_index +1 ]
	fmt.Println(number2[1:3])
	//
	// decision => if else
	//
	// var isTrue bool
	isTrue := true
	if isTrue { // if isTrue ==true
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
	cat := "dog"
	if cat == "cat" {
		log.Println("cat is", cat)
	} else {
		log.Println("cat is not", cat)
	}
	// cũng sử dụng &&  ||  như code C

	//
	// switch case
	//
	switchVar := "doasdasdg"
	switch switchVar {
	case "cat":
		fmt.Println("is a cat")
	case "dog":
		fmt.Println("is a dog")
	case "fish":
		fmt.Println("is a fish")
	case "bird":
		fmt.Println("is a bird")
	default:
		fmt.Println(switchVar, "is something else")
	}

	//
	//Loops
	//

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	animals := make(map[string]string)
	animals["dog"] = "Minh"
	animals["cat"] = "Chien"
	animals["bird"] = "Ha"
	animals["fly"] = "Cong"
	// in ra đi cùng thứ tự
	for i, animalName := range animals {
		fmt.Println(i, animalName)
	}
	// in ra ko kèm thứ tự
	for _, animalName := range animals {
		fmt.Println(animalName)
	}

	line := "Nguyen duy Chien"
	for i, lineNumber := range line {
		fmt.Println(i, ":", lineNumber)
	}

	var userSlices []User
	userSlices = append(userSlices, User{"Chien", "Nguyen", "Chien@gmail", 20})
	userSlices = append(userSlices, User{"Huyen", "Bui", "BTKH@gmail", 19})
	userSlices = append(userSlices, User{"MiNhon", "meo", "mew@mew", 1})
	for _, userName := range userSlices {
		fmt.Println(userName.FirstName)
	}

	//
	//interfaces
	//
	dog1 := Dog{"Mic", "Phu Quoc"}
	cat1 := Cat{"MiNhon", "Brown"}
	// dog1.Say()
	PrintInfo(&dog1)
	PrintInfo(&cat1)

	//
	//package
	//

	var myCallSOS helper.Sos
	myCallSOS.Element1 = "ao that day"
	fmt.Println(myCallSOS.Element1)

	//
	//channel
	//

	intChannel := make(chan int)
	defer close(intChannel)
	go CalculateValue(intChannel)
	num := <-intChannel
	fmt.Println("Random number is ", num)
}

const numPool = 1000

func CalculateValue(intChannel chan int) {
	randomNumber := helper.RandomNumber(numPool)
	intChannel <- randomNumber
}

// interfaces
// In order for something to implement an interface,
// it must implement the same function as the interface.
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Say(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Say() string {
	return "woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
func (c *Cat) Say() string {
	return "moew"
}
func (c *Cat) NumberOfLegs() int {
	return 4
}

func saySomething() string {
	return "nothing"
}
func changeThing(s *string) {
	new_color := "red"
	*s = new_color
}
