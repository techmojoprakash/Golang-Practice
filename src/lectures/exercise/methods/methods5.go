// Methods with non-struct receivers

// Prakash Notes: methods can create not only for struct types but also non struct types like type int, string etc...

// So far we have defined methods only on struct types.
// It is also possible to define methods on non-struct types,
// but there is a catch. To define a method on a type,
// the definition of the receiver type and the definition of
// the method should be present in the same package.
// So far, all the structs and the methods on structs we defined were all located in the same main package and hence they worked.

package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
