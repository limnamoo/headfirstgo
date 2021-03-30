package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file) // create scanner for file
	for scanner.Scan() { // read a line in file
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil { // if error while reading file, report error and exit program
		log.Fatal(scanner.Err())
	}
}
