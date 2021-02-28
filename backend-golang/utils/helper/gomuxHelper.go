package helper

import (
	"net/http"

	"github.com/gorilla/mux"
)

//GomuxPathVariable is a function for get path variable
func GomuxPathVariable(key string, request *http.Request) string {
	return mux.Vars(request)[key]
}
