package main

type Tracker struct {
	list map[string]string
}

func New(list map[string]string) *Tracker  {
	return &Tracker{list}
}

func (t * Tracker) Add()  {

}

func (t * Tracker) Dell()  {

}

func main() {
}

/*
type Sequence struct {
	message string
}

func New(newMessage string) *Sequence {
	return &Sequence{newMessage}
}

func (s *Sequence) FindMax() int {

	var maxCount int
	var count int
	arrRune := []rune(s.message)
	for _, runeValue := range arrRune {
		if string(runeValue) == "о" {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}
func combination(n int,m int) int {
	var m, n int
	fmt.Scan(&n)
	fmt.Scan(&m)
	println(combination(n, m))
	var (
		x = n
		y = m
		z = n-m
		xf = 1
		yf = 1
		zf = 1
	)

	for i := 1; i <=x ; i++ {
		println(i, xf)
		xf *= i
	}
	for i := 1; i <=y ; i++ {
		yf *= i
	}
	for i := 1; i <=z ; i++ {
		zf *= i
	}

	return xf / (yf * zf)
}

*/
