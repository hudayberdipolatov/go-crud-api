package routes

import (
	"github.com/gorilla/mux"
	"golang_crud_api/controllers"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", controllers.User{}.GetAllUser).Methods("GET")
	r.HandleFunc("/api/user/{id}", controllers.User{}.GetUser).Methods("GET")
	r.HandleFunc("/api/user/create", controllers.User{}.StoreUser).Methods("POST")
	r.HandleFunc("/api/user/update/{id}", controllers.User{}.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/delete/{id}", controllers.User{}.DeleteUser).Methods("DELETE")
	return r
}
