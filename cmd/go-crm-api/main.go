package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aralim11/go-crm-api/config"
	"github.com/aralim11/go-crm-api/infra/db"
	"github.com/aralim11/go-crm-api/internal/router"
)

func main() {
	// load config
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config Error!!")
	}

	// DB connection
	dbConn, err := db.NewConnection(cnf.Database)
	if err != nil {
		log.Fatal("Database Connection Error!!")
	}

	// DB migration
	err = db.MigrateDB(dbConn.DB, "file://infra/migrations", db.GetConnectionString(cnf.Database))
	if err != nil {
		log.Fatal("Database Migration Error!!")
	}

	// router
	mux := http.NewServeMux()
	router.RegisterModules(mux, dbConn)

	// start servers
	server := &http.Server{
		Addr:    cnf.Server.AppUrl + ":" + fmt.Sprintf("%s", cnf.Server.HTTPPort),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
