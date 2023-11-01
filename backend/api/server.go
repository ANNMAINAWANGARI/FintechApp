package api

import (
	"database/sql"
	"fmt"
	db "github/ANNMAINAWANGARI/FintechApp/db/sqlc"
	"github/ANNMAINAWANGARI/FintechApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	_ "github.com/lib/pq"
)


type Server struct {
	queries *db.Queries
	router  *gin.Engine
	config  *utils.Config
}


var tokenController *utils.JWTToken
var gValid = galidator.New().CustomMessages(
	galidator.Messages{
		"required": "this field is required",
	},
)

func NewServer(envPath string) *Server{

	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Couldn't load config: %v", err))
	}

	conn, err := sql.Open(config.DBdriver, config.DB_source_live+"?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}
	tokenController = utils.NewJWTToken(config)

	q:= db.New(conn)
	g := gin.Default()

	return &Server{
		queries: q,
		router:  g,
		config:  config,
		
	}
	

}

func (s *Server) Start(port int) {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Fingreat!"})
	})
    
	User{}.router(s)
	Auth{}.router(s)

	s.router.Run(fmt.Sprintf(":%v", port))
}