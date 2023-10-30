package main

import (
	"fmt"
	"golang_crud_api/models"
	"golang_crud_api/routes"
	"net/http"
)

func main() {

	fmt.Println("Server : https://localhost:8080")
	models.User{}.Migrate()
	http.ListenAndServe(":8080", routes.Routes())
}
