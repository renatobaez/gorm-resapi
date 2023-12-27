package domain

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	// Las clases u objetos que están en el dominio no debieran conocer nada de la lógica de la infraestructura
	// por lo que no debiera haber relación con las tablas de las bases de datos
	Title       string
	Description string
	Done        bool
	UserID      uint
}
