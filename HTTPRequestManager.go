
package main
import (

	"github.com/gin-gonic/gin"

)


func IsLoginUserExist(g *gin.Context) (bool, DBManager.DBUsers) {
    httpRequest := g.Request
    httpRequest.ParseForm() //これをやらないとリクエストが取れない。

    email := httpRequest.Form["email"][0]
    encryptedPassword := toHash(httpRequest.Form["password"][0])

    u := DBManager.GetLoginUser(email, encryptedPassword ,dbSession.Slave)

    return u.ID != 0, u
}