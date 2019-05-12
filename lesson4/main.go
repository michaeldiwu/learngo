package main

import (
	"fmt"
	"math"
)

type Student struct {
	Age  int
	Name string
}

func main() {

	/*
		bool
		byte
		int8 16 32 64 int8 int16
		uint8 16 32 64
		float8 16 32 64
		complex64 128 1+2i
		uintptr <== pointer
		string
		array
		struct
		function
		interface

		// different way to initialization
		map
		slice
		channel  <==
	*/

	s := "0Abc"
	fmt.Println([]byte(s)) // convert string s to byte array []byte --> byte array

	//b := false
	stu := new(Student) // malloc Student object and return address
	fmt.Println(*stu)   // print stu contents

	fmt.Println(math.MinInt8, math.MaxUint16)

	// map can not be created using new
	m := make(map[int]string)
	m[1] = "1"
	m[2] = "2"
	fmt.Println(m)

	//s1 := make([]int,0)
	//s1 = append(s1,12)
	//fmt.Println(s1)

	// channel is pipe
	c := make(chan int, 1) // 1 can push 1 item capacity
	c <- 1                 // push 1 into channel
	fmt.Println(<-c)       // get one item from channel

	/*p:=new(map[int]string)   // create a new map for {1: "test"} dict
	m1:=*p    // m1 is referring to map
	m1[1]="1"
	fmt.Println(m1)*/

	// type conversion
	a := 1
	b := uint8(a)
	fmt.Println(b)

	//// user defined type: give primitive types a different name for clarity
	type color int

	var red color
	red = 1
	var green color
	green = 2
	fmt.Println(red, green)

	/*red1 := int(1)
	red1 = red   //// error
	fmt.Println(red1)*/

	//reflect.Kind

}
