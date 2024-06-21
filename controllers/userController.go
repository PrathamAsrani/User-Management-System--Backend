package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/PrathamAsrani/User-Management-System--Backend/modals"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (userController UserController) GetUser (wr http.ResponseWriter, re http.Request, par httprouter.Params){
	id := par.ByName("id");
	if !bson.IsObjectIdHex(id) {
		wr.WriteHeader(http.StatusNotFound)
	}
	userId := bson.ObjectIdHex(id);
	user := modals.userModal{}
	if err := userController.Session.DB("EMP").C("Users").FindId(userId).One(&user); err != nil {
		wr.WriteHeader(http.StatusNotFound)
		return 
	}
	user_json, err := json.Marshal(user);
	if err != nil {
		fmt.Println(err);
	}
	wr.Header().set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOk);
	fmt.Fprintf(wr, "%s\n", user_json);
}


func (userController UserController) CreateUser (wr http.ResponseWriter, re http.Request, _ httprouter.Params){
	user := modals.userModal{}
	json.NewDecoder(re.Body).Decode(&user);
	user.Id = bson.NewObjectId();
	userController.session.DB("EMP").C("Users").Insert(user);

	user_json, err := json.Marshal(user);
	if err != nil {
		fmt.Println(err);
	}
	wr.Header().Set("Content-Type", "application/json");
	wr.WriteHeader(http.StatusCreated);
	fmt.Fprintf(wr, "%s\n", user_json);
}


func (userController UserController) DeleteUser (wr http.ResponseWriter, re http.Request, par httprouter.Params){
	id := par.ByName("id");
	if !bson.IsObjectIdHex(id) {
		wr.WriteHeader(http.StatusNotFound)
	}
	userId := bson.ObjectIdHex(id);
	if err := userController.Session.DB("EMP").C("Users").RemoveId(userId); err != nil {
		wr.WriteHeader(http.StatusNotFound)
		return 
	}
	wr.WriteHeader(http.StatusOk);
	fmt.Fprintf(wr, "User deleted\n", userId, "\n");
}