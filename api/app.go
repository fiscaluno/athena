package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/fiscaluno/athena/api/handler"
	"github.com/fiscaluno/athena/pkg/course"
	"github.com/fiscaluno/athena/pkg/detailedreview"
	"github.com/fiscaluno/athena/pkg/institution"
	"github.com/fiscaluno/athena/pkg/middleware"
	"github.com/fiscaluno/athena/pkg/review"
	"github.com/fiscaluno/athena/pkg/user"
	"github.com/fiscaluno/pandorabox"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// Start ...
func Start() {

	r := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//review
	reviewRepo := review.NewInmemRepository()
	reviewService := review.NewService(reviewRepo)
	handler.MakeReviewHandlers(r, *n, reviewService)

	// institution
	institutionRepo := institution.NewAPIRepository("http://aiolia.herokuapp.com", "/institution")
	institutionService := institution.NewService(institutionRepo)
	handler.MakeInstitutionHandlers(r, *n, institutionService)

	// course
	courseRepo := course.NewAPIRepository("http://shaka-course.herokuapp.com", "/courses")
	courseService := course.NewService(courseRepo)
	handler.MakeCourseHandlers(r, *n, courseService)

	// user
	userRepo := user.NewInmemRepository()
	userService := user.NewService(userRepo)
	handler.MakeUserHandlers(r, *n, userService)

	// detailed_review
	detailedReviewRepo := detailedreview.NewInmemRepository()
	detailedReviewService := detailedreview.NewService(detailedReviewRepo)
	handler.MakeDetailedReviewHandlers(r, *n, detailedReviewService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	port := pandorabox.GetOSEnvironment("PORT", "5001")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	logger.Println("Listen on port:" + port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
