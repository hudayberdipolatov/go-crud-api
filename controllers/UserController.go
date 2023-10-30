package controllers

import (
	"encoding/json"
	"fmt"
	"golang_crud_api/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct{}

func (user User) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})
	data["message"] = "Golang CRUD All users list"
	data["users"] = models.User{}.GetAllPost()
	json.NewEncoder(w).Encode(data)
}

func (user User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["message"] = "Golang CRUD GET user"
	data["user"] = models.User{}.GetPost(id)

	json.NewEncoder(w).Encode(data)
}

func (user User) StoreUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	data := make(map[string]string)
	json.Unmarshal(body, &data)

	username := data["username"]
	email := data["email"]
	fullname := data["fullname"]
	age := data["age"]

	models.User{
		Username: username,
		Email:    email,
		FullName: fullname,
		Age:      age,
	}.Create()
	fmt.Fprintf(w, "Taze User doredildi")
}

func (user User) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	user_update := models.User{}.GetPost(id)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	update := make(map[string]string)
	json.Unmarshal(body, &update)

	username := update["username"]
	email := update["email"]
	fullname := update["fullname"]
	age := update["age"]
	user_update.Updates(models.User{
		Username: username,
		Email:    email,
		FullName: fullname,
		Age:      age,
	})
	data := make(map[string]interface{})
	data["message"] = "Golang CRUD Update user"
	data["user"] = models.User{}.GetPost(id)
	json.Marshal(data)
	json.NewEncoder(w).Encode(data)
}

func (user User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	user_delete := models.User{}.GetPost(id)
	user_delete.Delete()
	fmt.Fprintf(w, "User Pozuldy!!!")
}
