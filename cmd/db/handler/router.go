package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ehs/cmd/db/domain"
	"github.com/ehs/cmd/db/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {

	db, client := ConnectDB()

	d := Dbhandlers{Service: service.Newdbuserservice(domain.NewDbHttppuser(db, client))}
	fmt.Println(d)
	r := mux.NewRouter()

	r.HandleFunc("/createuser", d.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/getuser/{field}/{id}", d.GetUserByFilter).Methods(http.MethodGet)
	r.HandleFunc("/create/room", d.CreateRoom).Methods(http.MethodPost)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}
	http.ListenAndServe(":"+port, r)

}
func init() {
	db, client := ConnectDB()
	d := Dbhandlers{Service: service.Newdbuserservice(domain.NewDbHttppuser(db, client))}
	fmt.Print(d)

}
func ConnectDB() (*mongo.Collection, *mongo.Client) {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("Details").Collection("Student_user")
	return collection, client

}
