package data

import (
	"encoding/json"
	"io"
)

type Moviment struct {
	Id    string  `json:"id"`
	Value float32 `json:"value"`
}

type Moviments []*Moviment

func (m *Moviments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func GetMoviment() Moviments {
	return movimentList
}

var movimentList = []*Moviment{
	{
		Id:    "1",
		Value: 2.45,
	},
	{
		Id:    "2",
		Value: 3.50,
	},
}
