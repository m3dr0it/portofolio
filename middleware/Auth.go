package middleware

import (
	"encoding/json"
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
				Data:    new(map[string]string),
			})

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
			return
		}

		authToken := strings.Replace(authTokenBearer, "Bearer ", "", -1)
		userInfo, err := service.ValidateJWT(authToken)

		if err != nil || userInfo.Username == "" {
			response, _ := json.Marshal(model.BaseResponse{
				Message: "Unauthorized",
				Data:    new(map[string]string),
			})

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
