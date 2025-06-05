package main

import (
	"net/http"
	"os"

	"github.com/grasieliw/GWStore/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback para local
	}

	http.ListenAndServe(":"+port, nil)
}
