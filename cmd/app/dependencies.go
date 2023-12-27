package app

import (
	"log"

	"github.com/renatobaez/gorm-resapi/internal/application"
	"github.com/renatobaez/gorm-resapi/internal/infrastructure/controller"
	"github.com/renatobaez/gorm-resapi/internal/infrastructure/db"
	"github.com/renatobaez/gorm-resapi/internal/infrastructure/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handlers struct {
	userController *controller.UserController
}

func BuildDependencies() *Handlers {
	var DSN = "host=localhost user=postgres password=123456 dbname=gorm port=5434"
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
	db.AutoMigrate(DB)

	// Aquí se hace la inversión de dependencias para cumplir con el principio de SOLID (D) por medio de la inyección de dependencias de cada servicio en la capa de más afuera.

	// Declarar repositories
	userRepository := repository.NewUserRepository(DB)

	// Declarar services
	userService := application.NewUserService(userRepository)

	// Declarar controller
	userController := controller.NewUserController(userService)

	return &Handlers{userController}
}
