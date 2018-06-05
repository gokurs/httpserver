package main

func main() {

}

// sum ist eine einfache Funktion mit zwei Integers als
// Input
func sum(a, b int) int {
	return a + b
}

// sumN summiert beliebig viele Integer zu einer Summe
func sumN(numbers ...int) int {
	var sum int
	for _, nr := range numbers {
		sum = sum + nr
	}
	return sum
}
