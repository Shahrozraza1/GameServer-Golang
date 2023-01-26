package api

import (
	"GameServer/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/** Function where all the api end points are created **/

func StartApiServer(ip, port string) {
	log.Println("Starting Gaming Server")
	router := mux.NewRouter()
	router.HandleFunc("/addlevel", handlers.PostAddALevelHandler).Methods("POST")
	router.HandleFunc("/getlevelbyid/{ID}", handlers.GetLevelByIDHandler).Methods("GET")
	router.HandleFunc("/updatelevel", handlers.PutUpdateALevel).Methods("PUT")
	router.HandleFunc("/deletelevel/{ID}", handlers.DeleteLevelHandler).Methods("DELETE")
	log.Printf("Server started on Port :%s", port)
	http.ListenAndServe(ip+":"+port, router)
}
