package controller

import (
	"encoding/json"
	"net/http"
	"portofolio/model"
	"portofolio/service"
)

func GetArticles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		articles, err := service.GetArticles()

		if err != nil {
			errResponse, _ := json.Marshal(model.BaseResponse{
				Message: err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errResponse)
		}

		response, _ := json.Marshal(model.BaseResponse{
			Message: "success",
			Data:    articles,
		})

		w.Write(response)
	})
}
