package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		count  int
		count2 int
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ = strconv.Atoi(scanner.Text())
	var arr []string
	for i := 0; i < count; i++ {
		scanner.Scan()
		arr = append(arr, scanner.Text())
	}
	scanner.Scan()
	count2, _ = strconv.Atoi(scanner.Text())
	var arr2 []int
	for i := 0; i < count2; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		arr2 = append(arr2, val)
	}
	for _, val := range arr2 {
		fmt.Println(arr[val-1])
	}
}

/*
	var (
		count int
		start int
		end int
		sum int
	)
	fmt.Scan(&count)
	var arr []int
	for i := 0; i <count; i++ {
		var val int
		fmt.Scan(&val)
		arr = append(arr, val)
	}
	fmt.Scan(&start)
	fmt.Scan(&end)
	fmt.Println(arr)
	arrSum := arr[start-1:end]
	fmt.Println(arrSum)
	for _, val := range arrSum {
			sum += val
	}
	fmt.Println(sum)


		arr := [4]string {"Да", "дА", "ДА", "да"}
		var str string
		fmt.Scan(&str)
		// r := []rune(str)
		var b bool
		for _, val := range arr {
			fmt.Println(val)
			if val == r[0] || val == len(r) {
		}
		for _, val := range str {
		fmt.Println(val)
		}
		if b {
		fmt.Println("СОГЛАСЕН")
		}else{
		fmt.Println("НЕ СОГЛАСЕН")
		}


		var countService, countMy int
		fmt.Scan(&countService)
		fmt.Scan(&countMy)
		var arrayService []string
		var arrayMy []string
		scanner := bufio.NewScanner(os.Stdin)
		for i := 0; i < countService; i ++ {
			scanner.Scan()
			text := scanner.Text()
			arrayService = append(arrayService, text)
		}
		for i := 0; i < countMy; i ++ {
			scanner.Scan()
			text := scanner.Text()
			arrayMy = append(arrayMy, text)
		}
		for _, val := range arrayMy {
			var exist bool
			for _, val2 := range arrayService {
				if !exist {
					exist = val == val2
				}
			}
			if exist {
				fmt.Println("ЕСТЬ")
			}else{
				fmt.Println("НЕТ В НАЛИЧИИ")
			}
		}

	var countService, countMy int
	fmt.Scan(&countService)
	fmt.Scan(&countMy)
	var arrayService []string
	var arrayMy []string
	for i := 0; i < countService; i ++ {
		var text string
		fmt.Scanln(&text)
		arrayService = append(arrayService, text)
	}
	for i := 0; i < countMy; i ++ {
		var text string
		fmt.Scanln(&text)
		arrayMy = append(arrayMy, text)
	}
	fmt.Println(arrayService, arrayMy)
*/
/*for _, val := range arrayMy {
	var exist bool
	for _, val2 := range arrayService {
		if !exist {
			exist = val == val2
		}
	}
	if exist {
		fmt.Println("ЕСТЬ")
	}else{
		fmt.Println("НЕТ В НАЛИЧИИ")
	}
}*/
/*
		var t, j, sum int
		fmt.Scan(&t)
		for i := 1; i <= t; i ++ {
			fmt.Scan(&j)
			if i % 2 == 0 {
				sum -= j
			} else {
				sum += j
			}
		}
		fmt.Println(sum)

		var t, j, sum int
		fmt.Scan(&t)
		for i := 1; i <= t; i ++ {
			fmt.Scan(&j)
			if i % 2 == 0 {
				sum -= j
			} else {
				sum += j
			}
		}
		fmt.Println(sum)
		for i, j := 0, 1; i <= t; i ++ {
			var str string
			for k := 1; k <= j; k++ {

				if k > i {
					str += strconv.Itoa(k-i)
				} else {
					str += strconv.Itoa(k)
				}
			}
			fmt.Println(str)
			j++
		}
	var g = -t
	for i := g; i <= t; i ++ {
		fmt.Printf("Квадрат числа %d равен %d\n", int(i) , int(math.Pow(i, 2)))
	}*/
/*arr := [5]string{"a", "f"}*/
