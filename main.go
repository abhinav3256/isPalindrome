package main

import "fmt"

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
}
func isPalindrome(x int) bool {
	temp := x
	revrese := 0

	for x > 0 {
		rem := x % 10
		revrese = revrese*10 + rem
		x = x / 10
	}

	return revrese == temp

}
