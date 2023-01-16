package routes

import (
	"log"
	"net/http"

	"github.com/CarlosGenuino/api-rest-go/controllers"
	"github.com/CarlosGenuino/api-rest-go/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.Personalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8500", r))
}
