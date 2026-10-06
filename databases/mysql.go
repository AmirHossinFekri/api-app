package databases

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)


type Params struct {
	User , Password , Host , Port ,Name string
}

func NewMySql(p Params ) (*sql.DB , error){
	cfg := mysql.NewConfig()
	cfg.User = p.User
	cfg.Passwd = p.Password
	cfg.Net = "tcp"
	cfg.Addr = p.Host + ":" + p.Port
	cfg.DBName = p.Name
	cfg.ParseTime = true
	cfg.Params = map[string]string{"charset":"utf8mb4"}


	db , err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil , fmt.Errorf("open: %w" ,err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5* time.Minute)


	if err := db.Ping(); err != nil{
		return nil, fmt.Errorf("ping: %w" , err)
	}

	return db,nil
}