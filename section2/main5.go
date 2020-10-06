package main

import "fmt"

//empty interfacwe
func tryEmptyInterface(value interface{}) {
	fmt.Println(value)
}

//interfaces : car type
type car interface {
	drive()
	stop()
}

// Lambo : type of car
type Lambo struct {
	Name string
}

// Chevvy : type of car
type Chevvy struct {
	Name string
}

func (l Lambo) drive() {
	fmt.Println(l.Name, " Lambo driving")
}

func (l Chevvy) drive() {
	fmt.Println(l.Name, " Chevvy driving")
}

func (l Lambo) stop() {
	fmt.Println(l.Name, " Lambo stopped")
}

func (l Chevvy) stop() {
	fmt.Println(l.Name, " chevvy stopped")
}

func instruct(c car) {
	fmt.Println(c)
	c.drive()
	c.stop()
}

// func main() {
// 	l := Lambo{"Lamborghini"}
// 	c := Chevvy{"Alto"}
// 	instruct(l)
// 	instruct(c)

// 	tryEmptyInterface(10)
// 	tryEmptyInterface("Swati")
// 	myMap := make(map[string]interface{})
// 	myMap["name"] = "swati"
// 	myMap["age"] = 10
// 	fmt.Println(myMap)
// }
