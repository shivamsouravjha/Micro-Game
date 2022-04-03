package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/shivamsouravjha/Micro-Game/UserService/config"

	_ "github.com/go-sql-driver/mysql"
)

var Dbmap = initDb()

func initDb() *gorp.DbMap {
	connection := config.Get().DBUserName + ":" + config.Get().DBPassword + "@tcp(" + config.Get().DBHostReader + ":" + config.Get().DBPort + ")/" + config.Get().DBName
	db, err := sql.Open(config.Get().DATABASE, connection)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	fmt.Println("connected")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
