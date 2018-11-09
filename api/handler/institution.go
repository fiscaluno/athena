package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/fiscaluno/athena/pkg/entity"
	"github.com/fiscaluno/athena/pkg/institution"
	"github.com/gorilla/mux"
)

func institutionIndex(service institution.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var data []*entity.Institution
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
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		resp := entity.HTTPResp{
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func institutionAdd(service institution.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var b *entity.Institution
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		b.ID, err = service.Store(b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
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
			w.Write([]byte(err.Error()))
			return
		}
	})
}

func institutionFind(service institution.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.Find(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		resp := entity.HTTPResp{
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func institutionDelete(service institution.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		err := service.Delete(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

//MakeInstitutionHandlers make url handlers
func MakeInstitutionHandlers(r *mux.Router, n negroni.Negroni, service institution.UseCase) {
	r.Handle("/v1/institutions", n.With(
		negroni.Wrap(institutionIndex(service)),
	)).Methods("GET", "OPTIONS").Name("institutionIndex")

	r.Handle("/v1/institutions", n.With(
		negroni.Wrap(institutionAdd(service)),
	)).Methods("POST", "OPTIONS").Name("institutionAdd")

	r.Handle("/v1/institutions/{id}", n.With(
		negroni.Wrap(institutionFind(service)),
	)).Methods("GET", "OPTIONS").Name("institutionFind")

	r.Handle("/v1/institutions/{id}", n.With(
		negroni.Wrap(institutionDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("institutionDelete")
}
