package main

import (
	"log"
	"net"
	"net/http"

	"github.com/amirhossinfekri/wiscky/config"
	"github.com/amirhossinfekri/wiscky/databases"
	"github.com/amirhossinfekri/wiscky/middleware"
	"github.com/amirhossinfekri/wiscky/router"
	"github.com/joho/godotenv"
)



func main(){

	_ = godotenv.Load()

	cfg := config.Load()


	db,err := databases.NewMySql(databases.Params{
		User:			cfg.DBUser,
		Password: cfg.DBPassword,
		Host: 		cfg.DBHost,
		Port:			cfg.DBPort,
		Name:     cfg.DBName,
	})


	if err != nil {
		log.Fatalf("database: %v" ,err)
	}

	defer db.Close()


	r :=router.New()

	handler := middleware.Apply(r,middleware.Logging,middleware.Recover)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	log.Printf("server is running on %s", ln.Addr())
	log.Fatal(http.Serve(ln, handler))
}