package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type dbConnection struct {
	db  *sql.DB
	config Config
}

func (d *dbConnection) iniDbConnection(){

	d.config = Config{
		Host: "localhost",
		Port: "5432",
		Name: "laundry_db",
		User: "postgres",
		Password: "postgres",
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", "postgres",d.config.User,d.config.Password,d.config.Host,d.config.Port,d.config.Name)

	db, err := sql.Open("postgres",dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err) 
	}
	d.db = db	
}

func GetConnection() *sql.DB{
	dbConn := dbConnection{}
	dbConn.iniDbConnection()
	return dbConn.db
}