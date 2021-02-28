package middlewares

import (
	"encoding/json"
	"log"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/helper"
	"net/http"
	"strings"
)

//AuthorizationToken is a function for token jwt security for an API
func AuthorizationToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getToken := r.Header.Get("Authorization")
		splitBearerToken := strings.Split(getToken, "Bearer ")
		if len(splitBearerToken) <= 1 {
			response := models.Response{
				Meta: models.MetaComponents{
					Status:  "Fail",
					Code:    401,
					Message: "No token key, cant access API",
					Records: 0,
				},
				Data: nil,
			}
			byteOfResponse, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write(byteOfResponse)
			log.Println(" | No token key")
			helper.LogApp("No token key")
			return
		}

		validity := helper.ValidateToken(splitBearerToken[1])
		if validity {
			next.ServeHTTP(w, r)
		} else {
			response := models.Response{
				Meta: models.MetaComponents{
					Status:  "Fail",
					Code:    401,
					Message: "Token expired or invalid token key, cant access API",
					Records: 0,
				},
				Data: nil,
			}
			byteOfResponse, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write(byteOfResponse)
			log.Println(" | Token expired or invalid token key")
			helper.LogApp("Token expired or invalid token key")
		}
	})
}
