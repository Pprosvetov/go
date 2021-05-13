package main

import "fmt"

func main() {
	var str string
	var count int
	var maxCount int
	fmt.Scan(&str)
	strR := []rune(str)

	for i, _ := range strR {
		if i > len(strR) {
			break
		}
		if string(strR[i]) == "о" {
			count++
			if len(strR) == i || string(strR[i + 1]) != "о" {
				if(maxCount < count) {
					maxCount = count
					count = 0
				}
			}
		}
	}
	fmt.Println(maxCount)
}

/*
	var f string
	fmt.Scan(&f)
	fr := []rune(f)
	i := 0

	for {
		if i % 2 == 0 {
			fr = fr[2: len(fr)]
		} else {
			fr = fr[0: len(fr) - 2]
		}
		i++
		fmt.Println(string(fr))
		if len(fr) <= 2 {
			break
		}
	}

	var f string
	fmt.Scan(&f)
	fr := []rune(f)
	fmt.Println(string(fr))
	i := 0

	for {
		if len(fr) <= 2 {
			break
		}
		if i % 2 == 0 {
			fr = fr[2: len(fr)]
		} else {
			fr = fr[0: len(fr) - 2]
		}
		i++
		fmt.Println(string(fr))
		if len(fr) <= 2 {
			break
		}
	}

	var str string
	var count int
	var maxCount int
	fmt.Scan(&str)
	strR := []rune(str)

	for i, _ := range strR {
		if i > len(strR) {
			break
		}
		if string(strR[i]) == "о" {
			count++
			if string(strR[i + 1]) != "о" {
				if(maxCount < count) {
					maxCount = count
					count = 0
				}
			}
		}
	}
	fmt.Println(maxCount)
	var f string
	fmt.Scan(&f)
	fr := []rune(f)
	for {
		var l string
		var few rune
		fmt.Scan(&l)
		lr := []rune(l)
		if string(fr[len(fr)-1]) == "ь" {
			few = fr[len(fr)-2]
		} else {
			few = fr[len(fr)-1]
		}

		if few == lr[0] {
			fr = lr
		} else {
			fmt.Println(l)
			break
		}
	}


	arr := [4]string{"Да", "дА", "ДА", "да"}
	var str string
	fmt.Scan(&str)
	r := []rune(str)
	var b bool
	f, l := string(r[0]), string(r[len(r)-1])
	for _, val := range arr {
		if val == f+l || val == l+f {
			b = true
		}
	}
	if b {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}

*/
