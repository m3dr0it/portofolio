package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func AddArticle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		defer r.Body.Close()
		var article model.Article
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
		}

		err = json.Unmarshal(body, &article)

		if err != nil {
			errResponse, _ := json.Marshal(model.BaseResponse{
				Message: err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errResponse)
		}

		isCreatedErr := service.AddArticle(article)

		if isCreatedErr != nil {
			errResponse, _ := json.Marshal(model.BaseResponse{
				Message: err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errResponse)
		}

		response, _ := json.Marshal(model.BaseResponse{
			Message: "success",
		})

		w.Write(response)
	})
}
