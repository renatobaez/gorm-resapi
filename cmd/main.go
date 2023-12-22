package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatobaez/gorm-resapi/cmd/app"
)

func main() {

	r := mux.NewRouter()

	application := app.BuildDependencies()
	app.LoadRoutes(r, application)
	// No hay problema con implementar las rutas en esta parte, pero hay que intentar dejarlas en la capa de infraestructura,
	// o en una instancia que nos permita generar la inversión de dependencias para que se cumplan los principios de SOLID.

	http.ListenAndServe(":3000", r)
}
