// // why do we have methods when we can write the same program using functions. There are a couple of reasons for this. Let's look at them one by one.

// // Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types are a way to achieve behaviour similar to classes. Methods allow a logical grouping of behaviour related to a type similar to classes. In the above sample program, all behaviours related to the Employee type can be grouped by creating methods using Employee receiver type. For example, we can add methods like calculatePension, calculateLeaves and so on.

// // Methods with the same name can be defined on different types whereas functions with the same names are not allowed. Let's assume that we have a Square and Circle structure. It's possible to define a method named Area on both Square and Circle. This is done in the program below.

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Rectangle struct {
// 	length int
// 	width  int
// }

// type Circle struct {
// 	radius float64
// }

// func (r Rectangle) Area() int {
// 	return r.length * r.width
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func main() {
// 	r := Rectangle{
// 		length: 10,
// 		width:  5,
// 	}
// 	fmt.Printf("Area of rectangle %d\n", r.Area())
// 	c := Circle{
// 		radius: 12,
// 	}
// 	fmt.Printf("Area of circle %f", c.Area())
// }

// // ----------done---------------------
