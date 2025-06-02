package routes

import (
	"GWStore/controllers"
	"net/http"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)

}
