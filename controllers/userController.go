package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PrathamAsrani/User-Management-System--Backend/modals"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (userController UserController) GetUser(wr http.ResponseWriter, re *http.Request, par httprouter.Params) {
	id := par.ByName("id")
	if !bson.IsObjectIdHex(id) {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	userId := bson.ObjectIdHex(id)
	user := modals.UserModal{}
	if err := userController.session.DB("EMP").C("Users").FindId(userId).One(&user); err != nil {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	fmt.Fprintf(wr, "%s\n", userJson)
}

func (userController UserController) CreateUser(wr http.ResponseWriter, re *http.Request, _ httprouter.Params) {
	user := modals.UserModal{}
	json.NewDecoder(re.Body).Decode(&user)
	user.Id = bson.NewObjectId()
	userController.session.DB("EMP").C("Users").Insert(user)

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	fmt.Fprintf(wr, "%s\n", userJson)
}

func (userController UserController) DeleteUser(wr http.ResponseWriter, re *http.Request, par httprouter.Params) {
	id := par.ByName("id")
	if !bson.IsObjectIdHex(id) {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	userId := bson.ObjectIdHex(id)
	if err := userController.session.DB("EMP").C("Users").RemoveId(userId); err != nil {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	wr.WriteHeader(http.StatusOK)
	fmt.Fprintf(wr, "User deleted\n")
}
