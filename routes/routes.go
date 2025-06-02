package routes

import (
	"net/http"

	"github.com/grasieliw/GWStore/controllers"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)

}
