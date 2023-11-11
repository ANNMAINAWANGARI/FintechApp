package api

import (
	"context"
	"database/sql"
	db "github/ANNMAINAWANGARI/FintechApp/db/sqlc"
	"github/ANNMAINAWANGARI/FintechApp/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	server *Server
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users",AuthenticatedMiddleware())
	serverGroup.GET("", u.listUsers)
	serverGroup.PATCH("/username", u.updateUsername)
	serverGroup.GET("/me", u.getLoggedInUser)

}

func (u *User) listUsers(c *gin.Context)  {
	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  10,
	}

	users, err := u.server.queries.ListUsers(context.Background(), arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    c.JSON(http.StatusOK, users)
	newUsers := []UserResponse{}

	for _, v := range users {
		n := UserResponse{}.toUserResponse(&v)
		newUsers = append(newUsers, *n)
	}

	c.JSON(http.StatusOK, newUsers)
}
func (u UserResponse) toUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *User) getLoggedInUser(c *gin.Context){
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}
	
	
	user, err := u.server.queries.GetUserByID(context.Background(), userId)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to access resources"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserResponse{}.toUserResponse(&user))
}

type UpdateUsernameType struct {
	Username string `json:"username" binding:"required"`
}


func (u *User) updateUsername(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	var userInfo UpdateUsernameType

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateUsernameParams{
		ID: userId,
		Username: sql.NullString{
			String: userInfo.Username,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	}

	user, err := u.server.queries.UpdateUsername(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{}.toUserResponse(&user))
}