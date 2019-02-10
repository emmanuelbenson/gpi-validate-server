package prevemployment

import (
	"encoding/json"
	"net/http"
)

// New adds a new Previous Employment Check request record
func New(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Previous Employment")
}

// Show displays a specific Previous Employment Check request record
func Show(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Display a specific Prev Emp request record")
}
