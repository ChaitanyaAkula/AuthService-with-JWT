package utils

import(
	"encoding/json"
	"net/http"
	
)
func ResponseJson(w http.ResponseWriter,data interface{}){
	w.Header().Set("Content-Type","apllication/json")
	json.NewEncoder(w).Encode(data)
}