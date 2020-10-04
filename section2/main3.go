package main

import "fmt"

//call by reference
func swap(m1 *int, m2 *int) { //swap(&m1, &m2)
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func pointer() {
	m := 2
	ptr := &m
	fmt.Println(*ptr)
	//ptr -> prints address ( & - address operator)
	//*ptr -> prints value at that address ( * - dereferencing operator)
}

// func main() {
// 	pointer()
// 	m1 := 3
// 	m2 := 4
// 	swap(&m1, &m2)
// 	fmt.Println(m1, m2)
// }
