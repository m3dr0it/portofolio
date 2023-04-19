package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"portofolio/model"
	"portofolio/service"
	"strings"
)

func ValidateToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authTokenBearer := r.Header.Get("Authorization")
		w.Header().Add("Content-Type", "application/json")

		if authTokenBearer == "" {
			response, _ := json.Marshal(model.BaseResponse{
				Message: "Unauthorized",
			})

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
			return
		}

		authToken := strings.Replace(authTokenBearer, "Bearer ", "", -1)
		fmt.Println(authToken)
		isValid, err := service.ValidateJWT(authToken)

		if err != nil || !isValid {
			response, _ := json.Marshal(model.BaseResponse{
				Message: "Unauthorized",
			})

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
