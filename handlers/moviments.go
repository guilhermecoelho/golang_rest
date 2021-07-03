package handlers

import (
	"gorest/data"
	"gorest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var moviment = models.Moviment{}

func GetMoviment(resp http.ResponseWriter, r *http.Request) {

	lp := data.GetMoviment()

	err := lp.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func GetMovimentById(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}

	lp, errorData := data.GetMovimentById(id)
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := lp.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func PutMoviment(resp http.ResponseWriter, r *http.Request) {

	errDecode := moviment.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PutMoviment(&moviment)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func PostMoviment(resp http.ResponseWriter, r *http.Request) {

	errDecode := moviment.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PostMoviment(&moviment)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func DeleteMoviment(resp http.ResponseWriter, r *http.Request) {

	errDecode := moviment.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	errorProcess := data.DeleteMoviment(&moviment)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
}
