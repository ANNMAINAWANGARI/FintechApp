package api

import (
	"context"
	"database/sql"
	"github/ANNMAINAWANGARI/FintechApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type Auth struct {
	server *Server
}

func (a Auth) router(server *Server) {
	a.server = server

	serverGroup := server.router.Group("/auth")
	serverGroup.POST("/login", a.login)

}



func (a Auth) login(c *gin.Context) {
	user := new(UserParams)
	eViewer := gValid.Validator(UserParams{})
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.HandleError(err, c, eViewer)})
		return
	}

	dbUser, err := a.server.queries.GetUserByEmail(context.Background(), user.Email)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := utils.VerifyPassword(user.Password, dbUser.HashedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	token, err := tokenController.CreateToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}