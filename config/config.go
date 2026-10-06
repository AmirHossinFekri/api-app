package config

import (
	"fmt"
	"log"
	"os"
)


type Config struct {
	DBUser 			string
	DBPassword  string
	DBHost 			string
	DBPort 			string 
	DBName 			string
}

func Load() Config{
	c :=Config{
		DBUser: mustGet("DB_USER"),
		DBPassword: mustGet("DB_PASSWORD"),
		DBHost: getOr("DB_HOST","127.0.0.1"),
		DBPort: getOr("DB_PORT" , "3306"),
		DBName: mustGet("DB_NAME"),
	}
	return c
}




func (c Config) DSN() string{
	return fmt.Sprintf(	"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)
}

func mustGet(key string) string{
	v := os.Getenv(key)
	
	if v == ""{
		log.Fatalf("missing required env : %s",key)
	}

	return v
}

func getOr(key string , replace string) string {
	v := os.Getenv(key)
	if v == ""{
		return replace
	}
	return v
}