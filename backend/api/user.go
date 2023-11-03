package api

import (
	"context"
	"database/sql"
	db "github/ANNMAINAWANGARI/FintechApp/db/sqlc"
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
	//Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users",AuthenticatedMiddleware())
	serverGroup.GET("", u.listUsers)
	
	serverGroup.POST("/me", u.getLoggedInUser)

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
		//Username:  user.Username.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
// func (u *User) createUser(c *gin.Context) {
// 	var user UserParams
// 	//single line error check
// 	if err :=c.ShouldBindJSON(&user); err!=nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	hashedPassword,err:=utils.GenerateHashPassword(user.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	arg := db.CreateUserParams{
// 		Email:          user.Email,
// 		HashedPassword: hashedPassword,
// 	}
// 	newUser, err := u.server.queries.CreateUser(context.Background(), arg)
// 	if err != nil {
// 		if pgErr, ok := err.(*pq.Error); ok {
// 			if pgErr.Code == "23505" {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, UserResponse{}.toUserResponse(&newUser))
// }
func (u *User) getLoggedInUser(c *gin.Context){
	value, exists := c.Get("user_id")
	if !exists{
		c.JSON(http.StatusUnauthorized,gin.H{"message":"Not authorized to access resource"})
        return
	}
	userId,ok :=value.(int64)
	if !ok{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Encountered an issue"})
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