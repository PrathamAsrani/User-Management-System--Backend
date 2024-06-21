package main

import (
    "log"
    "net/http"

    "github.com/PrathamAsrani/User-Management-System--Backend/controllers"
    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2"
)

func main() {
    log.Println("Starting the application...")

    // Create Router Instance
    router := httprouter.New()

    // Controllers
    userController := controllers.NewUserController(getSession())

    // Create REST API
    router.GET("/user/:id", userController.GetUser)
    router.POST("/user", userController.CreateUser)
    router.DELETE("/user/:id", userController.DeleteUser)

    // Server connect
    log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getSession() *mgo.Session {
    log.Println("Attempting to connect to MongoDB at mongodb://localhost:27017...")
    session, err := mgo.Dial("mongodb://localhost:27017/")
    if err != nil {
        log.Fatalf("Error while creating session: %v", err)
    }
    log.Println("Successfully connected to MongoDB")
    return session
}
