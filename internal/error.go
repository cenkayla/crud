package internal

import (
	"encoding/json"
	"net/http"
)

func SendErr(w http.ResponseWriter, err error, codes ...int) {
	if err == nil {
		return
	}
	if len(codes) < 1 {
		codes = append(codes, http.StatusInternalServerError)
	}
	response, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(response), codes[0])
}
