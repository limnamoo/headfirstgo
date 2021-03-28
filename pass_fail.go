package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // ParseFloat을 사용하기 위해 strconv패키지를 가져온다.
	"strings" // TrimSpace를 사용하기 위해 strings패키지를 가져온다.
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // 입력 문자열에서 줄바꿈 문자를 제거한다.
	grade, err := strconv.ParseFloat(input, 64) // 문자열을 float64 값으로 변환한다.
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
