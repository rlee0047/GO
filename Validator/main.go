package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type uSer struct {
	ID        string `validate:"omitempty,uuid"`
	Firstname string `validate:"omitempty"` //Ether omitempty or required
	Lastname  string `validate:"omitempty"` //Ether omitempty or required
	Username  string `validate:"required,email"`
	Password  string `validate:"required,gte=10"`
	Type      string `validate:"isdefault"`
}

func main() {
	user := uSer{
		Firstname: "Ryan",
		Lastname:  "Lee",
		Username:  "test@gmail.com",
		Password:  "1234567890",
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Welcome,", user.Firstname, user.Lastname)
	}
}
