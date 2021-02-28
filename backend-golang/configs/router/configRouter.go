package router

import (
	"fmt"
	"log"
	"mekarTechnicalTest/utils/helper"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//ConfigRouter is a function for create new router server and port
func ConfigRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

//RunServer is a function for running server and port
func RunServer(router *mux.Router) {
	server := helper.GetKeyValue("SERVER", "localhost")
	port := helper.GetKeyValue("PORT", "8001")
	serverAndPort := fmt.Sprintf("%v:%v", server, port)
	log.Printf(" | Server run : %v", server)
	log.Printf(" | Port run : %v", port)
	helper.LogApp("Server run : " + server)
	helper.LogApp("Port run : " + port)
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(serverAndPort, handler)
	helper.FatalError(err, "Router Error")
}
