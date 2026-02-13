package main

import (
	"conceitos_avancados/foo"
	"encoding/json"
	"fmt"
)

type MinhaString string

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`

	foo.Foo
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func UpdateName(u *User, newName string) {
	u.Name = newName
}

func main() {

	user := User{Name: "Orlando", ID: 1}
	user.UpdateName("Joaquim")
	fmt.Println(user)
	UpdateName(&user, "Pedro")
	fmt.Println(user)
	user.Bar()

	res, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
