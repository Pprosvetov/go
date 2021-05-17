package main

func main() {
	/*var num int
	var isLight bool
	var text string
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		if num % i == 0 {
			text = text + fmt.Stringer(i)
		}
	}*/

	/*var t float64
	fmt.Scan(&t)
	var g = -t
	for i := g; i <= t; i ++ {
		fmt.Printf("Квадрат числа %d равен %d\n", int(i) , int(math.Pow(i, 2)))
	}*/

	/*
		var (
			min float64 = 100
			max float64 = 140
			minF float64
			maxF float64
			count int
		)
		for {
			var pulse float64
			fmt.Scan(&pulse)
			if pulse < 0 {
				fmt.Println(count)
				fmt.Printf("%.1f %.1f\n", minF, maxF)
				return
			}
			if pulse >= min && pulse <= max {
				if pulse < minF || minF == 0 {
					minF = pulse
				}
				if pulse > maxF || maxF == 0 {
					maxF = pulse
				}
				count++
				continue
			}
		}
			var a int64
			fmt.Scan(&a)
			fmt.Printf("%b\n", a)
		for {
			var pas1, pas2 string
			fmt.Scan(&pas1)
			fmt.Scan(&pas2)
			if utf8.RuneCountInString(pas1) < 8 {
				fmt.Println("Слишком короткий пароль!")
				continue
			}
			if strings.Contains(pas1, "123") || strings.Contains(pas1, "qwe") {
				fmt.Println("Слишком простой пароль!")
				continue
			}
			if pas1 != pas2 {
				fmt.Println("Введенные пароли различаются!")
				continue
			}
			fmt.Println("Ну наконец-то!" )
			break
		}


		for {
			var text string
			fmt.Scan(&text)
			fmt.Println(0)
			fmt.Println(reflect.TypeOf(text))
			if text == "" {
				fmt.Println(1)
				return
			}
			fmt.Println(text)
		}*/
}
