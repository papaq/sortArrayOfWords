package iao

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func LineToConsole(str string) {
	fmt.Println(str)
}

func LinesToConsole(str []string) {
	for _, line := range str {
		LineToConsole(line)
	}
}

func WordInfoToConsole(word string, distance int) {
	fmt.Printf("%15v%5v\n", word, distance)
}

func ReadFile(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	
	countries := strings.Split(string(data), "\n")[:]
	return countries[0:len(countries) - 1]
}

