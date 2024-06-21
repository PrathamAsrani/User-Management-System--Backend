package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"github.com/PrathamAsrani/User-Management-System--Backend/controllers"
)

func main() {
	fmt.Println("hello")

	// create Router Instance
	router := httprouter.New()

	// controllers
	userController := controllers.NewUserController(getSession())

	// create Rest API
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	// server connect
	http.ListenAndServe("localhost:8080", router)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Print("Error while creating session")
		panic(err)
	}
	return session
}
