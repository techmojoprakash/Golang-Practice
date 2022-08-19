//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

type Rectangle struct {
	lengthX int
	widthX  int
}

// perimeter of a rectangle
func getRectPeremeter(lengthX, widthX int) int {
	return 2 * lengthX * widthX
}

// area of rectangle
func getRectArea(lengthX, widthX int) int {
	return lengthX * widthX
}

// func main() {
// 	myRect := Rectangle{2, 3}
// 	fmt.Println(getRectPeremeter(myRect.lengthX, myRect.widthX))
// 	fmt.Println(getRectArea(myRect.lengthX, myRect.widthX))

// }
