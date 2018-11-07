package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/fiscaluno/athena/pkg/detailedreview"
	"github.com/fiscaluno/athena/pkg/entity"
	"github.com/gorilla/mux"
)

func detailedreviewIndex(service detailedreview.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading detailedreviews"
		var data []*entity.DetailedReview
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.FindAll()
		default:
			data, err = service.Search(name)
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

func detailedreviewAdd(service detailedreview.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding detailedreview"

		vars := mux.Vars(r)
		reviewID := vars["review_id"]

		var b *entity.DetailedReview
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		u, err := strconv.ParseUint(reviewID, 10, 64)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		b.ReviewID = uint(u)
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

func detailedreviewFind(service detailedreview.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading detailedreview"
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

func detailedreviewDelete(service detailedreview.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing detailedreview"
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

func detailedreviewAverage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading detailedreviews"

		type average struct {
			ReviewType string  `json:"review_type"`
			Rate       float64 `json:"rate"`
		}

		var data []*average
		var err error
		// name := r.URL.Query().Get("name")

		data = []*average{
			{
				ReviewType: "Infrastructure",
				Rate:       4.5,
			},
			{
				ReviewType: "Professors",
				Rate:       2.5,
			},
			{
				ReviewType: "Classes",
				Rate:       1.5,
			},
			{
				ReviewType: "Website",
				Rate:       3.5,
			},
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

//MakeDetailedReviewHandlers make url handlers
func MakeDetailedReviewHandlers(r *mux.Router, n negroni.Negroni, service detailedreview.UseCase) {
	r.Handle("/v1/reviews/{review_id:[0-9]+}/details", n.With(
		negroni.Wrap(detailedreviewIndex(service)),
	)).Methods("GET", "OPTIONS").Name("detailedreviewIndex")

	r.Handle("/v1/reviews/{review_id:[0-9]+}/details", n.With(
		negroni.Wrap(detailedreviewAdd(service)),
	)).Methods("POST", "OPTIONS").Name("detailedreviewAdd")

	r.Handle("/v1/reviews/{review_id:[0-9]+}/details/{id}", n.With(
		negroni.Wrap(detailedreviewFind(service)),
	)).Methods("GET", "OPTIONS").Name("detailedreviewFind")

	r.Handle("/v1/reviews/{review_id:[0-9]+}/details/{id}", n.With(
		negroni.Wrap(detailedreviewDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("detailedreviewDelete")

	r.Handle("/v1/reviews/details/average", n.With(
		negroni.Wrap(detailedreviewAverage()),
	)).Methods("GET", "OPTIONS").Name("detailedreviewAverage")
}
