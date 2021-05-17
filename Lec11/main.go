package main

import "fmt"

func main() {
	var (
		w int
		d int
		s string
	)
	fmt.Scan(&w)
	fmt.Scan(&d)
	fmt.Scan(&s)
	for i := 1; i <= w; i++ {
		var str string
		for j := 1; j <= d; j++ {
			if j == 1 || j == d || i == 1 || i == w {
				str = str + s
			} else {
				str = str + string(' ')
			}
			//fmt.Println(w, d, i, j, str)
		}
		fmt.Println(str)
	}
}

/*
	var (
		b int = 20
		k int = 10
		t int = 5
		sum int
		count int
	)
	fmt.Scan(&sum)
	fmt.Scan(&count)

	var (
		w int
		d int
		s string
	)
	fmt.Scan(&w)
	fmt.Scan(&d)
	fmt.Scan(&s)
	for i := 1; i <= w; i++ {
		var str string
		for j := 1; j <= d; j++ {
			if j == 1 || j == d ||  i == 1 || i == w{
				str = str + s
			}else {
				str = str + string(' ')
			}
			//fmt.Println(w, d, i, j, str)
		}
		fmt.Println(str)
	}

	var w int
	fmt.Scan(&w)
	for i := 1; i <= w; i++ {
		var str string
		for k := 1; k <= w-i; k++ {
			str = str + string(' ')
		}
		for j := 1; j <= i + i-1; j++ {
			str = str + string('@')
		}
		fmt.Println(str)
	}

	fmt.Scan(&w)
	if(w % 2 == 0){
		fmt.Println("YES")
	}else {
		fmt.Println("NO")
	}


*/
