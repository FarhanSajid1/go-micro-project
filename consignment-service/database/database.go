package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

var MONGO_USER = getEnv("MONGOUSER", "jamesbond")
var MONGO_PASSWORD = getEnv("MONGOPASSWORD", "password")
var MONGO_PORT = getEnv("MONGOPORT", "27017")
var MONGO_HOST = getEnv("MONGOHOST", "localhost")

func ConnectDB() *mgo.Session {
	str := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		MONGO_USER, MONGO_PASSWORD,MONGO_HOST, MONGO_PORT)
	session, err := mgo.Dial(str)
	if err != nil {
		log.Printf("could not connect %v", err)
	}

	return session
}
