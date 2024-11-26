// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"simple_api/routes"
	"simple_api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// Add custom recovery middleware
func customRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				// get error from panic
				err, ok := rvr.(error)
				if !ok {
					err = fmt.Errorf("internal server error: %v", rvr)
				}
				utils.ErrorResponse(w,
					http.StatusInternalServerError,
					fmt.Sprintf("%v", rvr))
				panic(err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set up zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Create a new router
	r := chi.NewRouter()

	// Use middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customRecoverer)

	// Register routes
	routes.RegisterRoutes(r)

	// Start server
	log.Info().Msgf("Starting server on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
