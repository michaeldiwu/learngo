package lesson1

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 1000
	var y float32
	var z string

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//var a, b int
	//var t, s = 1000, "sad"
	/*var (
		a int
		b string
		c float32
	)*/
	r := 1 // new variable with value
	fmt.Println(r)
	fmt.Println(reflect.TypeOf(r))

	left, right := 1, 2
	left, right = right, left
	fmt.Println(left, right)
}
