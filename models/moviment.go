package models

import (
	"encoding/json"
	"io"
)

type Moviment struct {
	Id    int     `json:"id"`
	Value float32 `json:"value"`
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
