package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/fiscaluno/athena/pkg/entity"
	"github.com/fiscaluno/athena/pkg/review"
	"github.com/gorilla/mux"
)

func reviewIndex(service review.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading reviews"
		var data []*entity.Review
		var err error
		query := r.URL.Query().Encode()
		switch {
		case query == "":
			data, err = service.FindAll()
		default:
			data, err = service.Search(query)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		resp := entity.HTTPResp{
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func reviewAdd(service review.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding review"
		var b *entity.Review
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		b.ID, err = service.Store(b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		resp := entity.HTTPResp{
			Code:   http.StatusCreated,
			Result: b,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func reviewFind(service review.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading review"
		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.Find(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		resp := entity.HTTPResp{
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func reviewDelete(service review.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing review"
		vars := mux.Vars(r)
		id := vars["id"]
		err := service.Delete(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeReviewHandlers make url handlers
func MakeReviewHandlers(r *mux.Router, n negroni.Negroni, service review.UseCase) {
	r.Handle("/v1/reviews", n.With(
		negroni.Wrap(reviewIndex(service)),
	)).Methods("GET", "OPTIONS").Name("reviewIndex")

	r.Handle("/v1/reviews", n.With(
		negroni.Wrap(reviewAdd(service)),
	)).Methods("POST", "OPTIONS").Name("reviewAdd")

	r.Handle("/v1/reviews/{id:[0-9]+}", n.With(
		negroni.Wrap(reviewFind(service)),
	)).Methods("GET", "OPTIONS").Name("reviewFind")

	r.Handle("/v1/reviews/{id:[0-9]+}", n.With(
		negroni.Wrap(reviewDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("reviewDelete")
}
