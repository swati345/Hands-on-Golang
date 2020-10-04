package main

import "fmt"

//structure -> implements encapsulation

//Car : stores general info
type Car struct {
	Name  string
	Age   int
	Model int
}

func (c Car) print() {
	fmt.Println(c)
}

func (c Car) getName() string {
	return c.Name
}
func structures() {
	c := Car{"santro", 2, 12}
	//fmt.Println(c)
	c.print()
	fmt.Println(c.getName())
	// var c Car

}

// func main() {
// 	structures()
// }
