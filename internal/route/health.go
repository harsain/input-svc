package route

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/harsain/input-svc/internal/model"
	"time"
)

//HealthCheck returns a health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	healthCheck := model.HealthCheck{
		Status: "OK",
		CreatedAt: time.Now().String(),
	}
	json.NewEncoder(w).Encode(healthCheck)
}
