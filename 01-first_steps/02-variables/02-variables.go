package main

import "fmt"

func main() {
	
	/* 
		var var_name type
		- It allow us to create an object of static type
		- Also it asssign a default value to the variable
	*/
	fmt.Println("Creating a numer")
	var number int
	fmt.Println(number)
	number = 25
	fmt.Println(number)
	number = 40
	fmt.Println(number)

	/*
		var_name := 50
		- This form allow us to create and assign an object, in this case an int
	*/

	/*
		In Go we can not assign crated variables with a diferent type of self type
		var name string
		name = 15 <- This will throw us an error because name is an string type variable
		name = 'John' <- This is correct
	*/

	/*
		Double assignation
		var_name_1, var_name_2 = 10,20
	*/
	fmt.Println("Direct creation and assignation")
	number_A := 5
	number_B := 6
	fmt.Println(number_A, number_B)

	fmt.Println("Double assignation")
	number_A, number_B = 10, 15
	fmt.Println(number_A, number_B)

	/*
		We can swap the values just using:
		var_name_1, var_name_2 = var_name_2, var_name_1
	*/
	fmt.Println("Swaping variables")
	fmt.Println(number_A, number_B)
	number_A, number_B = number_B, number_A
	fmt.Println(number_A, number_B)

	/*
		Also we can create and assignate directly two or more variables
	*/

	fmt.Println("Create and assignate two or more variables")
	number_C, number_D := 10, 10
	fmt.Println(number_C, number_D)

	fmt.Println("Four created and assigned variables")
	a, b, c, d := 1,2,3,4
	fmt.Println(a, b, c, d)

	// Note: We can create and assignate two variables, with one new name

	number_D, number_E := number_A, 20
	fmt.Println(number_D, number_E)
}