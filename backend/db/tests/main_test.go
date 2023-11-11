package db_test

import (
	"database/sql"
	db "github/ANNMAINAWANGARI/FintechApp/db/sqlc"
	"github/ANNMAINAWANGARI/FintechApp/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)


var testQuery *db.Store
const sslmode = "?sslmode=disable"

func TestMain(m *testing.M){
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load env config", err)
	}
	conn, err := sql.Open(config.DBdriver, config.DB_source_live+sslmode)
	if err != nil{
		log.Fatal("Could not connect to db",err)
	}
	testQuery = db.NewStore(conn)
	os.Exit(m.Run())
}