package routes

import (
	"net/http"

	"github.com/grasieliw/GWStore/controllers"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)

}
