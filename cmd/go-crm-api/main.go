package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aralim11/go-crm-api/config"
	"github.com/aralim11/go-crm-api/infra/db"
	"github.com/aralim11/go-crm-api/internal/user"
)

func main() {
	// load config
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config Error!!")
	}

	// DB connection
	db, err := db.NewConnection(cnf.Database)
	if err != nil {
		log.Fatal("Database Connection Error!!")
	}

	// handlers
	userRepository := user.NewUserRepo(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	// router
	router := http.NewServeMux()
	user.RegisterRoutes(router, userHandler)

	// start servers
	server := &http.Server{
		Addr:    cnf.Server.AppUrl + ":" + fmt.Sprintf("%s", cnf.Server.HTTPPort),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
