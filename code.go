package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	
	r := mux.NewRouter()
	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(getAllUsers)

	http.ListenAndServe(":8080",r)
}
func getAllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("worksss....")
}
