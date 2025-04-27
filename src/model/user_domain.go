package model

import (
	"encoding/json"
	"fmt"
)

type userDomain struct {
	Id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJsonValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetID() string {
	return ud.Id
}

func (ud *userDomain) SetID(id string) {
	ud.Id = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}