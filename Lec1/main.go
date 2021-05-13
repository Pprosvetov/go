package main

import "fmt"

func main() {
	var (
		var1 int
		var2 int
	)
	fmt.Scan(&var1)
	fmt.Scan(&var2)
	fmt.Println((var1 + var2) * (var1 + var2))
}

/*
	var (
		drink string
		meal string
		subMeal string
		time string
	)
	fmt.Scan(&drink)
	fmt.Scan(&meal)
	fmt.Scan(&subMeal)
	fmt.Scan(&time)
	//fmt.Println("I wanna some '" + drink + "', my favorite meal - '" + meal + "'. Give me also '" + subMeal + "'. I will come soon... '" + time + "'")
	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%v'\n", drink, meal, subMeal, time)

	var (
		lName string
		fName string
		age string

	)

	fmt.Scan(&lName)
	fmt.Scan(&fName)
	fmt.Scan(&age)
	fmt.Println("Имя: " + lName + " , Фамилия: " + fName + " , Возраст: " + age + " . Студент BPS");

	var (
		word1 = "имя"
		word2 = "твое"
		word3 = "мне"
		word4 = "знакомо"
	)
	fmt.Println(word4 + " " + word3 + " " + word2 + " " + word1 + " ")
	fmt.Println(word3 + " " + word1 + " " + word4 + " " + word2 + " ")
	fmt.Println(word2 + " " + word4 + " " + word1 + " " + word3 + " ")


*/
