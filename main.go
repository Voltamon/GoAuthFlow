package main

import (
	"log"

	data "github.com/Voltamon/GoAuthFlow/API/Data"
	service "github.com/Voltamon/GoAuthFlow/API/Service"
	transport "github.com/Voltamon/GoAuthFlow/API/Transport"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := service.InitDB("./API/Data/user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = service.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := data.NewUserRepo(db)
	serviceSvc := service.NewAuthService(userRepo)
	handler := transport.NewHandler(serviceSvc)

	router := gin.New()
	transport.SetupRoutes(router, handler)

	log.Println("Server starting on :8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
