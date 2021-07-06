package model

import (
	"regexp"
	"strings"
)

type User struct {
	Email    string
	Password string
}

type Users []User

var rgx = regexp.MustCompile(`\((.*?)\)`)

func (u *User) GetPasswordAndName(value string) {
	data := rgx.FindStringSubmatch(value)
	dataUser := strings.Split(data[1], ",")
	u.Email = strings.Trim(dataUser[0], " ")
	u.Password = strings.Trim(dataUser[1], " ")
}

