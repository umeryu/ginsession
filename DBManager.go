package main

import (
	"github.com/gocraft/dbr"

)



type DBManager struct {
	GetLoginUser()
	type DBUsers struct {
		ID int
		Username string
	}
}

func GetLoginUser(email string, password string, sess *dbr.Session) DBUsers {
   	var u DBUsers
	sess.Select("*").From("users").Where("email = ? AND password = ?", email, password).Load(&u)
	return u
}
