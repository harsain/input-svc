package route

import "github.com/gorilla/mux"

//RegisterRoutes registers all the routes of the app
func RegisterRoutes(r *mux.Router) {
	r.StrictSlash(true)

	// register routes
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/sensordata", GetSensorData).Methods("GET")
}