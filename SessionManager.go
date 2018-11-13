package main

import (
	"github.com/gin-gonic/gin"
)


type DBManager struct {
	
}

func ClearSession(g *gin.Context) {
    session := sessions.Default(g)
    session.Clear()
    session.Save()
    println("Session clear")
}

func Login(g *gin.Context, user DBManager.DBUsers) {
    session := sessions.Default(g)
    session.Set("alive", true)
    session.Set("userID", user.ID)
    session.Set("nickName", user.NickName)
    session.Save()
}