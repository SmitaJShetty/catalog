package common

import (
	"encoding/json"
	"net/http"
	"strings"
)

//IfExists if item exists in the array, return true else false
func IfExists(itemToBeFound string, itemArray []string) bool {
	found := false

	for _, item := range itemArray {
		if strings.EqualFold(item, itemToBeFound) {
			found = true
		}
	}

	return found
}

//IfIntExists if item exists in the array, return true else false
func IfIntExists(itemToBeFound int, itemArray []int) bool {
	found := false

	for _, item := range itemArray {
		if item == itemToBeFound {
			found = true
		}
	}

	return found
}

//SendResult sends result over http response
func SendResult(w http.ResponseWriter, r *http.Request, resultJSON []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

//SendErrorResponse sends error response
func SendErrorResponse(w http.ResponseWriter, r *http.Request, appErr *AppError) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json, _ := json.Marshal(appErr)
	w.Write(json)
}
