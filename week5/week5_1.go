package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScore = strings.TrimSpace(inputScore)
	score, err := strconv.ParseInt(inputScore, 32)

	if inputScore >= 90 {
		grade := "A grade"
	} else {
		grade := "under A grade"
	}
	fmt.Println(intScore)
	fmt.Println("You will get", grade) //undefined: grade
	fmt.Println(inputScore)
}
