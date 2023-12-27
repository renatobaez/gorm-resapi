package db

import (
	"github.com/renatobaez/gorm-resapi/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

// La instancia con la base de datos debiera hacerse en el momento que instanciamos el resto de las capas, para poder ocuparla en la inyección de dependencias
// por lo que tenerla instanciada desde el main nos quitará esa posibilidad, y nos acoplaremos a la infraestructura.
func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(models.User{}, models.Task{})
}
