package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const fName = "test.txt"

type statement struct {
	State string `json:"state"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile(fName)
	check(err)
	dats := []byte(dat)

	var word []statement

	err = json.Unmarshal(dats, &word)
	check(err)

	for i, v := range word {
		fmt.Println("\nStatement Number", i+1)
		fmt.Println(v.State)
	}

}
