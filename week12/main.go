package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0} // slice literal

	s[4] = 99
	s[2] = 11

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}
}
