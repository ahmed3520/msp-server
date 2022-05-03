package router

import (
	"github.com/ataha3520/msp-server/middleware"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/speakers", middleware.GetAllSpeakers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/addspeaker", middleware.AddSpeakerToEvent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/speaker/{id}", middleware.getSpecificSpeaker).Methods("GET", "OPTIONS")
	return router
}
