package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(getAllUsers)

	http.ListenAndServe(":8080",r)
}
func getAllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("worksss....")
}
