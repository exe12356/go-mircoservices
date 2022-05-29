package mongo

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	Dsn() string
	DbName() string
}

type config struct {
	User   string
	Pass   string
	Host   string
	Port   int
	dbName string
	dsn    string
}

func NewConfig() *config {
	var config config
	config.User = os.Getenv("MONGO_DB_USER")
	config.Pass = os.Getenv("MONGO_DB_PASS")
	config.Host = os.Getenv("MONGO_DB_HOST")
	config.dbName = os.Getenv("MONGO_DB_dbName")
	var err error
	config.Port, err = strconv.Atoi(os.Getenv("MONGO_DB_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	config.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", config.User, config.Pass, config.Host, config.Port, config.dbName)
	return &config
}

func (c *config) Dsn() string {
	return c.dsn
}
func (c *config) DbName() string {
	return c.dbName
}
