package controller

import (
	"encoding/json"
	"net/http"

	"github.com/codesleep-Beperfect/Bookstore/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Firstname string   `json:"firstname"`
	Lastname  string	`json:"lastname"`
	Email     string	`json:"email"`
}

var db *gorm.DB

func init(){
	utils.Connect()
	db = utils.Migrate()
	db.AutoMigrate(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user []User
	db.Find(&user)
	json.NewEncoder(w).Encode(user)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	if !checkifexist(params["id"]){
		json.NewEncoder(w).Encode("User ID Not Found!")
		return
	} 
	var user User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	if !checkifexist(params["id"]) {
		json.NewEncoder(w).Encode("User ID Not Found!")
		return
	}
	var user User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	if !checkifexist(params["id"]){
		json.NewEncoder(w).Encode("User ID Not Found!")
		return
	}
	var user User
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("User Deleted Successfully!!")
}

func checkifexist(id string) bool{
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		return false
	}
	return true
}