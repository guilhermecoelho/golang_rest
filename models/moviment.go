package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Moviment struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
}

type Moviments []*Moviment

func (m *Moviments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Moviment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Moviment) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
