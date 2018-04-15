package route

import (
	"net/http"
	"encoding/json"
	"github.com/harsain/input-svc/internal/db"
)

//GetSensorData returns the Sensor Data
func GetSensorData(w http.ResponseWriter, r *http.Request) {

	// use the DB to interact with DynamoDB
	db := db.GetDB()

	json.NewEncoder(w).Encode("Still alive")
}
