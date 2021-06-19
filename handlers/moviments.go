package handlers

import (
	"gorest/data"
	"net/http"
)

func GetMoviment(resp http.ResponseWriter, r *http.Request) {

	lp := data.GetMoviment()

	err := lp.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}
