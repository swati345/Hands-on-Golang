package main

import (
	"fmt"
	"strings"
)

/*
Atomic data types -
string
int
int32
int64
uint
uint32
uint64
*/

func intAndString() {

	var m int
	m = 2
	fmt.Println(m)

	//implicit declaration (:= --> declare and assign )
	l1 := 2
	l2 := 3
	fmt.Println(l1 + l2)

	var (
		m1 = 2
		m2 = 3
	)
	fmt.Println(m1 + m2)

	//typecasting
	var (
		i1 int32
		i2 int64
	)
	fmt.Println(int64(i1) + i2)

	//All strings are mutable
	var s string
	s = "My name is swati"
	fmt.Println(s)

	s1 := "Hey!"
	s1 = "Hello" //Cannot use := as it decalres and initialize and we cannot declae same variable again.
	fmt.Println(s1)

	//string datatype has many functions in strings package

	s2 := "My name"
	s3 := "name"
	fmt.Println(strings.Contains(s2, s3))
	fmt.Println(strings.Split(s2, " "))
}
